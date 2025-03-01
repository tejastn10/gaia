<p align="center">
  <img src="logo.svg" alt="Logo">
</p>

# Gaia 🌀

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/tejastn10/gaia?logo=go&logoColor=white)
[![Unit Tests](https://github.com/tejastn10/gaia/actions/workflows/unit-test.yml/badge.svg)](https://github.com/tejastn10/gaia/actions/workflows/unit-test.yml)
[![Release Workflow](https://github.com/tejastn10/gaia/actions/workflows/release.yml/badge.svg)](https://github.com/tejastn10/gaia/actions/workflows/release.yml)
![License](https://img.shields.io/badge/License-MIT-yellow?logo=open-source-initiative&logoColor=white)

Gaia is a Go-based CLI tool that allows you to modify the first commit of a Git repository easily. It ensures proper structure and provides user prompts for force pushing when necessary.

---

## Features 🌟

- **First Commit Modification**: Edit and replace the first commit message while maintaining commit integrity.
- **Interactive Rebase Automation**: Automates interactive rebase to allow easy commit modification.
- **Git Repository Checks**:
  - Verifies if the directory is a Git repository.
  - Ensures the working tree is clean before proceeding.
- **User Prompts for Force Push**: Asks for confirmation before force pushing changes.
- **Error Handling**: Handles rebase failures gracefully and provides recovery options.

---

## Getting Started

### Installation ⚙️

You can install Gaia using the following command:

```bash
curl -sSf https://raw.githubusercontent.com/tejastn10/gaia/main/scripts/install.sh | bash
```

This will download and install the latest version of Gaia.

### Uninstallation

To remove Gaia from your system, run:

```bash
curl -sSf https://raw.githubusercontent.com/tejastn10/gaia/main/scripts/uninstall.sh | bash
```

This will remove the installed binary from your system.

---

### Example Usage

```bash
$ gaia "Initial commit with updated message"

Successfully updated first commit: "Initial commit with updated message"
Do you want to force push the changes? (yes/no):
```

### Contributing 🤝

Contributions are welcome! Check out the [Contribution Guide](CONTRIBUTING.md) for details on how to get started. Feel free to open an issue or submit a pull request if you have ideas to enhance Gaia.

### To-Do ✅

- Add support for advanced commit manipulation.
- Provide additional safety checks before force pushing.
- Enhance logging and error messages.

---

## License 📜

This project is licensed under the MIT License. See the [LICENSE](LICENSE.md) file for details.

---

## Acknowledgments 🙌

- Named after **Gaia**, symbolizing stability and foundation.
- Built with ❤️ and Go.
