# Contributing to Gaia

Thank you for considering contributing to Gaia! Your contributions help improve the project and make it better for everyone. Follow the guidelines below to contribute effectively.

---

## 📌 Getting Started

1. **Fork the Repository**: Click the 'Fork' button on the top-right corner of the [repository](https://github.com/tejastn10/gaia) to create your own copy.
2. **Clone the Repository**:

    ```bash
    git clone https://github.com/tejastn10/gaia.git
    cd gaia
    ```

3. **Create a New Branch**:

    ```bash
    git checkout -b feature/your-feature-name
    ```

4. **Install Dependencies**:

    ```bash
    go mod tidy
    ```

---

### Project Structure 📂

```bash
gaia/
├── cmd/                  # CLI commands
│   ├── root.go           # Main CLI entry point
├── scripts/              # Scripts
│   ├── install.sh        # Install script
│   ├── uninstall.sh      # Un-Install script
├── tasks/                # Core logic for commit update
│   ├── commit.go         # Commit update functions
│   ├── repo.go           # Repository functions
├── go.mod                # Go module definition
├── go.sum                # Go module checksum
├── main.go               # Application entry point
├── README.md             # Project documentation
├── LICENSE.md            # Project license
```

---

## 🚀 Making Changes

1. **Implement Your Changes**:
   - Follow best practices and maintain consistency in the code.
   - Ensure your changes do not introduce breaking issues.

2. **Run and Test the Application**:

    ```bash
    go run main.go
    ```

    - Make sure everything runs as expected.

3. **Build the Project**:

    ```bash
    go build -o gaia
    ```

4. **Run the CLI Tool**:

    ```bash
    ./gaia --dir="/your/directory/path"
    ```

5. **Commit Your Changes**:

    ```bash
    git add .
    git commit -m "feat: Added feature XYZ"
    ```

6. **Push the Changes**:

    ```bash
    git push origin feature/your-feature-name
    ```

---

## ✅ Submitting a Pull Request

1. Navigate to the original repository: [Gaia](https://github.com/tejastn10/gaia).
2. Click on the **'New Pull Request'** button.
3. Select your fork and branch and compare it with the `main` branch.
4. Add a meaningful title and description for your changes.
5. Submit the pull request and wait for review.

---

## 🛠 Code Guidelines

- Use clear and descriptive commit messages.
- Follow the existing project structure.
- Write clean, readable, and maintainable code.
- Avoid unnecessary changes in unrelated files.

---

## 💬 Need Help?

If you have any questions or need clarification, feel free to open an issue or discuss it in the repository's [Issues](https://github.com/tejastn10/gaia/issues).

Happy coding! 🚀
