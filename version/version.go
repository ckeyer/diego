package version

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/Masterminds/semver"
)

const (
	timefmt = "2006-01-02T15:04:05Z-0700"
)

// -X git.funxdata.com/brisk/brisk/version.version=$(VERSION)
// -X git.funxdata.com/brisk/brisk/version.gitCommit=$(GIT_COMMIT)
// -X git.funxdata.com/brisk/brisk/version.buildAt=${date "+%Y-%m-%dT%H:%M:%SZ%z"}
var (
	version, gitCommit string
	buildAt            string

	v   semver.Version
	bAt time.Time
)

func init() {
	if version == "" {
		return
	}
	ver, err := semver.NewVersion(version)
	if err != nil {
		panic(fmt.Errorf("format version failed, %s", err))
	}
	v = *ver

	if gitCommit != "" {
		v, _ = ver.SetMetadata(gitCommit)
	}

	if buildAt != "" {
		t, err := time.Parse(timefmt, buildAt)
		if err != nil {
			panic(fmt.Errorf("format buildAt failed, %s", err))
		}
		bAt = t
	}
}

// GetVersion
func GetVersion() string {
	return fmt.Sprintf("%d.%d.%d", v.Major(), v.Minor(), v.Patch())
}

// GetSemver ...
func GetSemver() semver.Version {
	return v
}

// GetInt64
func GetInt64() int64 {
	return getInt64(v)
}

// getInt64	...
func getInt64(sv semver.Version) int64 {
	if sv.Patch() >= 10000 || sv.Minor() >= 100 {
		return 0
	}
	n := sv.Major()*100 + sv.Minor()
	n = n*10000 + sv.Patch()
	return int64(n)
}

// GetGitCommit
func GetGitCommit() string {
	return gitCommit
}

// GetCompleteVersion
func GetCompleteVersion() string {
	return v.String()
}

func GetBuildAt() time.Time {
	return bAt
}

func GetBuildAtString() string {
	return bAt.Format(timefmt)
}

func Print(w io.Writer) {
	if w == nil {
		w = os.Stdout
	}

	fmt.Fprintf(w, "Version:    %s\n", GetCompleteVersion())
	fmt.Fprintf(w, "Git Commit: %s\n", GetGitCommit())
	fmt.Fprintf(w, "Build At:   %s\n", GetBuildAtString())
}

// Map ...
func Map() map[string]string {
	return map[string]string{
		"version":    GetCompleteVersion(),
		"git_commit": GetGitCommit(),
		"build_at":   GetBuildAtString(),
	}
}

// IsSupport 是否是支持的版本
func IsSupport(ver string) bool {
	if strings.Contains(ver, "-") {
		ver = ver[strings.LastIndex(ver, "-")+1:]
	}
	inver, err := semver.NewVersion(ver)
	if err != nil {
		return false
	}
	return !inver.LessThan(&v)
}
