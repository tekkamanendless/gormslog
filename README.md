# gormslog
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/tekkamanendless/gormslog?label=version&logo=version&sort=semver)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/tekkamanendless/gormslog)](https://pkg.go.dev/github.com/tekkamanendless/gormslog)

This package provides a Gorm logger that uses `log/slog`.

# Usage
```
gormConfig := &gorm.Config{
	Logger: gormslog.New(),
}
db, err := gorm.Open(dialector, gormConfig)
```
