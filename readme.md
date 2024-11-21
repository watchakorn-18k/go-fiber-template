<div align="center">

# go-fiber-template

![img](https://camo.githubusercontent.com/41d946f07862663ea1040702d26c14cd0b1984f65002cdbaf61d57998426e247/68747470733a2f2f646f63732e676f66696265722e696f2f696d672f6c6f676f2d6461726b2e737667)

</div>

## Install golang package

```bash
go mod tidy
```

# Start APP

```sh
go run . || go run main.go
```

## Air

```sh
go install github.com/cosmtrek/air@latest
```

```sh
air
```

# Podman

```
podman build -t fiber-test .
```

```
podman run --rm -it -p 3000:3000 fiber-test
```
