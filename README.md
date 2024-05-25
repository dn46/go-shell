# Go-Shell

Go-Shell is a simple shell implemented in Go. It supports basic commands, navigation with `cd`, and exiting with `exit`.

## Features

- Execute any command available in the PATH
- Navigate directories with `cd`
- `cd` with no arguments returns to the home directory
- Exit the shell with `exit`
- Colored output for username, hostname, and current directory

## Usage

To run the shell, simply execute:

```bash
make run
```

## Code Structure

- `Shell`: The main shell struct that contains a reader to read input from the user.
- `Run`: The main loop of the shell that prints the prompt, reads input, and executes the input.
- `execInput`: A function that takes a string input, parses it, and executes the command.

## Acknowledgements

This project is inspired and based on `Writing a simple shell in Go` found on this [blog post](https://blog.init-io.net/post/2018/07-01-go-unix-shell/). Give love to the author for making it available.

## License

[MIT](https://choosealicense.com/licenses/mit/)