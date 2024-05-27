# Terraform Teeworlds

![golangci-lint](https://github.com/theobori/terraform-teeworlds/actions/workflows/lint.yml/badge.svg) ![build](https://github.com/theobori/terraform-teeworlds/actions/workflows/build.yml/badge.svg)

Entertaining Terraform chaos engineering, destroy resource by capturing Teeworlds flag.

## 📖 Build and run

For the build, you only need the following requirements:

- [Go](https://golang.org/doc/install) 1.22.3


Next to the Go application you could need the following requirements:
- [Teeworlds server](https://www.teeworlds.com/?page=downloads&id=14786) 0.7
  - With a econ server

## 🤝 Contribute

If you want to help the project, you can follow the guidelines in [CONTRIBUTING.md](./CONTRIBUTING.md).

## 💥 How to start chaos ?

There is a [`examples`](./examples) folder where you can find a terraform project.

You could run it with the following command from the project root folder.

```bash
terraform -chdir=./examples apply
```

You should also create a Teeworlds server (only tested on 0.7 version) with a similar configuration.

```bash
# Econ configuration
ec_port 7000
ec_password "hello_world"
ec_output_level 2

# Override server configuration
sv_register 0
sv_map ctf1
sv_gametype ctf
```

Now you can build and run the Go application, check the `-h` or `--help` flag if needed.
