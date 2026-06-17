Excellent. Here's a more realistic, senior-level Composite exercise without any code.

---

# Task: Company Organization Hierarchy

You're building an HR system.

## Components

There are two kinds of entities:

### Individual Employees (Leaf)

Examples:

* Software Engineer
* Designer
* Tester
* DevOps Engineer

They cannot contain anyone else.

---

### Managers (Composite)

Examples:

* Team Lead
* Engineering Manager
* CTO
* CEO

Managers can manage:

* Individual employees
* Other managers

---

## Hierarchy

Build the following organization:

```text
CEO
│
├── CTO
│     │
│     ├── Backend Lead
│     │      ├── Developer A
│     │      └── Developer B
│     │
│     └── Frontend Lead
│            ├── Developer C
│            └── Designer
│
└── HR Manager
       ├── Recruiter A
       └── Recruiter B
```

---

# Operations

Implement the ability to:

### 1. Print Organization

Calling:

```text
CEO.Show()
```

should recursively display the whole hierarchy.

---

### 2. Count Employees

Calling:

```text
CEO.CountEmployees()
```

should return the total number of people under the CEO.

Managers count as employees too.

---

### 3. Calculate Salaries

Each person has a salary.

Calling:

```text
CEO.TotalSalary()
```

should recursively compute the sum of salaries of everyone below.

---

### 4. Add Members

Managers should be able to add:

* Individual employees
* Other managers

---

### 5. Remove Members

Managers should be able to remove subordinates.

---

# Bonus Features

## Find Employee

Search for:

```text
"Developer C"
```

starting from CEO.

Should recursively traverse the hierarchy.

---

## Count Developers

Return how many developers exist in the entire organization.

---

## Average Salary

Compute average salary across the organization.

---

## Department Budget

Find total salaries under:

```text
CTO
```

without involving HR.

---

# Why this exercise is good

You'll naturally discover:

```text
Employee
↑
Manager
↑
CEO
```

where a manager contains:

```text
Employees
Managers
```

and those managers again contain:

```text
Employees
Managers
```

which is the essence of Composite:

> **A group and an individual are treated through the same interface.**

---

## Super Bonus (Closest to Real Go Backend)

Build a filesystem:

```text
root
│
├── home
│     ├── user
│     │      ├── notes.txt
│     │      └── pictures
│     │             ├── cat.jpg
│     │             └── dog.jpg
│
└── etc
      └── config.yaml
```

Support:

* Print tree
* Find file
* Calculate total size
* Count files
* Add/remove files and folders

This is probably the most famous and practical Composite pattern example and very close to how operating systems and cloud storage systems are modeled.
