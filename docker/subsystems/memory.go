package subsystems

type MemorySubSystem struct {

}

func (s *MemorySubSystem) Set(cgroupPath string, res *ResourceConfig) error {

	return nil
}

func (s *MemorySubSystem) Apply(cgroupPath string, pid int) error  {
	if subsysCgroupPath,err := GetCgroupPath {
		
	}
}
func (s *MemorySubSystem) Name() string  {
	return "memory"
}