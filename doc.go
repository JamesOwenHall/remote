// Package remote provides a server that listens on a Unix socket to execute
// custom commands.
//
// Usage
//
// Creating a server:
//     r, err := remote.New("/tmp/demo.sock")
//     if err != nil {
//         panic(err)
//     }
//
// Registering commands:
//     r.Commands["dev-mode"] = func(args []string) string {
//         if len(args) != 1 {
//             return "dev-mode commands takes exactly 1 argument: true or false"
//         }
//
//         switch args[0] {
//         case "true":
//             // ...
//             return "dev-mode is enabled"
//         case "false":
//             // ...
//             return "dev-mode is disabled"
//         default:
//             return "dev-mode command expects either true or false"
//         }
//     }
//
//     r.Commands["quit"] = func(args []string) string {
//         r.StopListening()
//         return ""
//     }
//
// Starting the server:
//     // ListenAndServe blocks, so you'll often want to start it as a new
//     // goroutine.
//     go r.ListenAndServe()
package remote
