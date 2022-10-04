package node

var nodeCmds = []string{
	"install",
	"uninstall",
	"ci",
}

func RunNpmCommand(inputCommand []string) []string {
	cmd := inputCommand[0]

	switch cmd {
	case
		"access",
		"adduser",
		"audit",
		"bin",
		"bugs",
		"build",
		"bundle",
		"cache",
		"ci",
		"completion",
		"config",
		"dedupe",
		"deprecate",
		"dist-tag",
		"docs",
		"doctor",
		"edit",
		"explore",
		"fund",
		"help",
		"help-search",
		"hook",
		"init",
		"install",
		"install-ci-test",
		"install-test",
		"link",
		"logout",
		"ls",
		"org",
		"outdated",
		"owner",
		"pack",
		"ping",
		"prefix",
		"profile",
		"prune",
		"publish",
		"rebuild",
		"repo",
		"restart",
		"root",
		"run-script",
		"search",
		"shrinkwrap",
		"star",
		"stars",
		"start",
		"stop",
		"team",
		"test",
		"token",
		"uninstall",
		"unpublish",
		"update",
		"version",
		"view",
		"whoami":
		return append([]string{"npm"}, inputCommand...)
	default:
		return inputCommand
	}
}
