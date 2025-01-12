# Go CI/CD Playground

## Usage

For now there are two behaviors:

1. If you push a new tag ([following standads](https://goreleaser.com/limitations/semver/)) it will automatically build and release. (example: [here](https://github.com/calmonr/cicd/releases))
2. If you send a pull request it will lint. (example: [here](https://github.com/calmonr/cicd/pull/1))

## Tools

> You can use [act](https://github.com/nektos/act) to run GitHub Actions locally.

- [x] [GitHub Actions](https://docs.github.com/pt/actions)
- [x] [GoReleaser](https://goreleaser.com/)
- [x] [golangci-lint](https://golangci-lint.run/)
- [x] [EditorConfig](https://editorconfig.org/)
- [ ] [Gitleaks](https://gitleaks.io/)
- [ ] [Cobra](https://cobra.dev/)
- [ ] [gRPC](https://grpc.io/)
  - [ ] [Buf](https://buf.build/)
- [ ] [Dagger](https://dagger.io/)
- [ ] [Task](https://taskfile.dev/)
- [ ] [UPX](https://github.com/upx/upx)

## Plans

- [ ] [CLI completion](https://carlosbecker.com/posts/golang-completions-cobra/)
- [ ] [Versioning](https://goreleaser.com/cookbooks/using-main.version)
- [x] [Docker images](https://goreleaser.com/customization/docker/)
  - [x] GitHub Container Registry
  - [x] Docker Hub
- [ ] Sign [checksums, artifacts](https://goreleaser.com/customization/sign/), [Docker images and manifests](https://goreleaser.com/customization/docker_sign/)
- [x] Tests

> Probably more to come.

## Tips

You should use [pre-commit](https://pre-commit.com/) to run linting and testing on your project, this way you can catch problems early.
> I'm not going to do that here (we want action :eyes:), but you should.

```yaml
---
repos:
  - repo: https://github.com/golangci/golangci-lint
    rev: v1.48.0
    hooks:
      - id: golangci-lint
default_install_hook_types: [pre-commit]
minimum_pre_commit_version: 2.20.0
```
