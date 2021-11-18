package user_agent

func (p *UserAgent) detectMicroMessenger(sections []section) {
	for _, s := range sections {
		if s.name == "MicroMessengerVersion" {
			p.microMessengerVersion = s.version
		}
	}
}
