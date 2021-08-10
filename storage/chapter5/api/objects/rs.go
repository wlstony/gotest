package objects

import (
	"fmt"
	"github.com/api/constant"
	"github.com/klauspost/reedsolomon"
	"io"
)

type RSPutStream struct {
	*encoder
}

func NewRSPutStream(dataServers []string, hash string, size int64) (*RSPutStream, error) {
	if len(dataServers) != constant.AllShards {
		return nil, fmt.Errorf("dataServers number mismatch require %d, now %d", constant.AllShards, len(dataServers))
	}
	perShard := (size + constant.DataShard - 1) / constant.DataShard
	writers := make([]io.Writer, constant.AllShards)
	var e error
	for i := range writers {
		writers[i], e = NewTempPutStream(dataServers[i], fmt.Sprintf("%s.%d", hash, i), perShard)
		if e != nil {
			return nil, e
		}
	}
	enc := NewEncoder(writers)
	return &RSPutStream{enc}, nil

}

func (s *RSPutStream) Commit(success bool) {
	s.Flush()
	for i := range s.writers {
		s.writers[i].(*TempPutStream).Commit(success)
	}
}

type encoder struct {
	writers []io.Writer
	enc     reedsolomon.Encoder
	cache   []byte
}

func NewEncoder(writers []io.Writer) *encoder {
	enc, _ := reedsolomon.New(constant.DataShard, constant.ParityShards)
	return &encoder{
		writers: writers,
		enc:     enc,
		cache:   nil,
	}
}

func (e *encoder) Write(p []byte) (n int, err error) {
	fmt.Println("encoder write: ", string(p))
	length := len(p)
	current := 0
	for length != 0 {
		next := constant.BlockSize - len(e.cache)
		if next > length {
			next = length
		}
		e.cache = append(e.cache, p[current:current+next]...)
		if len(e.cache) == constant.BlockSize {
			e.Flush()
		}
		current += next
		length -= next
	}
	return len(p), nil
}
func (e *encoder) Flush() {
	if len(e.cache) == 0 {
		return
	}
	shards, _ := e.enc.Split(e.cache)
	e.enc.Encode(shards)
	for i := range shards {
		e.writers[i].Write(shards[i])
	}
	e.cache = []byte{}
}

type decoder struct {
	readers   []io.Reader
	writers   []io.Writer
	enc       reedsolomon.Encoder
	size      int64
	cache     []byte
	cacheSize int
	total     int64
}

func NewDecoder(readers []io.Reader, writers []io.Writer, size int64) *decoder {
	enc, _ := reedsolomon.New(constant.DataShard, constant.ParityShards)
	return &decoder{
		readers:   readers,
		writers:   writers,
		enc:       enc,
		size:      size,
		cache:     nil,
		cacheSize: 0,
		total:     0,
	}
}
func (d *decoder) Read(p []byte) (n int, err error) {
	if d.cacheSize == 0 {
		e := d.getData()
		if e != nil {
			return 0, e
		}
	}
	length := len(p)
	if d.cacheSize < length {
		length = d.cacheSize
	}
	d.cacheSize -= length
	copy(p, d.cache[:length])
	d.cache = d.cache[length:]
	return length, nil
}
func (d *decoder) getData() error {
	if d.total == d.size {
		return io.EOF
	}
	shards := make([][]byte, constant.AllShards)
	repairIds := make([]int, 0)
	for i := range shards {
		if d.readers[i] == nil {
			repairIds = append(repairIds, i)
		} else {
			shards[i] = make([]byte, constant.BlockPerShard)
			n, e := io.ReadFull(d.readers[i], shards[i])
			if e != nil && e != io.EOF && e != io.ErrUnexpectedEOF {
				shards[i] = nil
			} else if n != constant.BlockPerShard {
				shards[i] = shards[i][:n]
			}
		}
	}
	e := d.enc.Reconstruct(shards)
	if e != nil {
		return e
	}
	for i := range repairIds {
		id := repairIds[i]
		d.writers[id].Write(shards[id])
	}
	for i := 0; i < constant.DataShard; i++ {
		shardSize := int64(len(shards[i]))
		if d.total+shardSize > d.size {
			shardSize -= d.total + shardSize - d.size
		}
		d.cache = append(d.cache, shards[i][:shardSize]...)
		d.cacheSize += int(shardSize)
		d.total += shardSize
	}
	return nil
}

type RSGetStream struct {
	*decoder
}

func NewRSGetStream(locateInfo map[int]string, dataServers []string, hash string, size int64) (*RSGetStream, error) {
	if len(locateInfo)+len(dataServers) != constant.AllShards {
		return nil, fmt.Errorf("dataServers number mismatch")
	}
	readers := make([]io.Reader, constant.AllShards)
	for i := 0; i < constant.AllShards; i++ {
		server := locateInfo[i]
		if server == "" {
			locateInfo[i] = dataServers[0]
			dataServers = dataServers[1:]
			continue
		}
		reader, e := NewGetStream(server, fmt.Sprintf("%s.%d", hash, i))
		if e == nil {
			readers[i] = reader
		}
	}
	writers := make([]io.Writer, constant.AllShards)
	perShard := (size + constant.DataShard - 1) / constant.DataShard
	var e error
	for i := range readers {
		if readers[i] == nil {
			writers[i], e = NewTempPutStream(locateInfo[i], fmt.Sprintf("%s.%d", hash, i), perShard)
			if e != nil {
				return nil, e
			}
		}
	}

	dec := NewDecoder(readers, writers, size)
	return &RSGetStream{dec}, nil
}

func (s *RSGetStream) Close() {
	for i := range s.writers {
		if s.writers[i] != nil {
			s.writers[i].(*TempPutStream).Commit(true)
		}
	}
}
