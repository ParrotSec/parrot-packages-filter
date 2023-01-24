package constants

const (
	PrefixName       = "Package: "
	PrefixDesc       = "Description: "
	PrefixVersion    = "Version: "
	PrefixMaintainer = "Maintainer: "
	PrefixArch       = "Architecture: "
	PrefixSection    = "Section: "
)

var Branch = [3]string{
	"contrib",
	"main",
	"non-free",
}

var Arch = [4]string{
	"amd64",
	"arm64",
	"armhf",
	"i386",
}
