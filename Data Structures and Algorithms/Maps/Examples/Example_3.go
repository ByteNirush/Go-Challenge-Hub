// Example 3: Grouping Data
// This example demonstrates how to use a map to group data based on a common attribute.

package main

import "fmt"

// User represents a user with a name and a role.
type User struct {
    Name string
    Role string
}

func main() {
    users := []User{
        {Name: "Alice", Role: "Admin"},
        {Name: "Bob", Role: "Editor"},
        {Name: "Charlie", Role: "Admin"},
        {Name: "David", Role: "Viewer"},
        {Name: "Eve", Role: "Editor"},
    }

    // Group users by role
    usersByRole := make(map[string][]User)

    for _, user := range users {
        usersByRole[user.Role] = append(usersByRole[user.Role], user)
    }

    // Print the grouped users
    for role, users := range usersByRole {
        fmt.Printf("Role: %s\n", role)
        for _, user := range users {
            fmt.Printf("  - %s\n", user.Name)
        }
    }
}
// In this example, we have a slice of User structs. We use a map to group the users by their roles. The keys of the map are the roles, and the values are slices of users with that role.

