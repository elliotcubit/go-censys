package decoded

type Redis struct {
	ArchBits            string                 `json:"arch_bits"`
	AuthResponse        string                 `json:"auth_response"`
	BuildId             string                 `json:"build_id"`
	Commands            string                 `json:"commands"`
	CommandsProcessed   uint64                 `json:"commands_processed"`
	ConnectionsReceived uint64                 `json:"connections_received"`
	GccVersion          string                 `json:"gcc_version"`
	GitSha1             string                 `json:"git_sha1"`
	InfoResponse        map[string]string      `json:"info_response"`
	Major               uint64                 `json:"major"`
	MemAllocator        string                 `json:"mem_allocator"`
	Minor               uint64                 `json:"minor"`
	Mode                string                 `json:"mode"`
	NonexistentResponse string                 `json:"nonexistent_response"`
	Os                  string                 `json:"os"`
	PatchLevel          uint64                 `json:"patch_level"`
	PingResponse        string                 `json:"ping_response"`
	QuitResponse        string                 `json:"quit_response"`
	RawCommandOutput    Redis_RawCommandOutput `json:"raw_command_output"`
	Uptime              uint64                 `json:"uptime"`
	UsedMemory          uint64                 `json:"used_memory"`
}

type Redis_RawCommandOutput struct {
	Output string `json:"output"`
}
