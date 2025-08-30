# Todo CLI


## **🎯 Goal**

Build a simple command-line application to manage tasks (add, list, complete). Tasks are stored locally in a JSON file.

The app should be **easy to use**, follow **Go idioms**, and be **testable**.

## **📌 Features**

### **1. Add a Task**

- Command:

```bash
todo add "Buy milk"
```

- Behavior:
    - Appends a new task to the task list.
    - Each task has:
        - ID (auto-increment int)
        - Description (string)
        - Completed (bool)
        - CreatedAt (timestamp)
    - Saves updated list to tasks.json.

### **2. List Tasks**

- Command:

```bash
todo list

```

- Behavior:
    - Shows all tasks, ordered by creation.
    - Output format example:

```bash
[ ] 1. Buy milk
[✓] 2. Write Go project
```

• Optional: Add flag for --completed or --pending.

### **3. Mark Task as Done**

- Command:

```bash
todo done 1
```

- Behavior:
    - Marks task with given ID as completed.
    - Saves changes back to tasks.json.
    - If ID doesn’t exist, show error.

### **4. Delete a Task**

- Command:

```bash
todo delete 2
```

- Behavior:
    - Removes task with given ID.
    - Saves changes back to tasks.json.

---

### **📂 Project Structure**

```bash
todo-cli/
  ├── main.go
  ├── cmd/
  │    ├── add.go
  │    ├── list.go
  │    ├── done.go
  │    ├── delete.go
  │    ├── add_test.go      # tests for add command
  │    └── list_test.go     # tests for list command, etc.
  ├── task/
  │    ├── task.go
  │    ├── store.go
  │    ├── task_test.go     # tests for Task struct behavior
  │    └── store_test.go    # tests for persistence (JSON read/write)
  ├── tasks.json
  ├── go.mod
```

---

## **✅ Acceptance Criteria**

- Commands work as described.
- Data persists between runs in tasks.json.
- Graceful error messages for invalid inputs (bad ID, empty description, etc.).
- Unit tests for:
    - Adding a task.
    - Marking a task done.
    - Deleting a task.
    - Loading/saving tasks.

---

## **🔧 Stretch Goals (Optional)**

- Add --format=json flag to output machine-readable JSON.
- Add --sort=created|completed.
- Support editing a task description.
- Replace JSON with SQLite for persistence.
