// JavaScript example for typing practice
class TaskManager {
    constructor(name) {
        this.name = name;
        this.tasks = new Map();
        this.nextId = 1;
    }

    // Add a new task
    addTask(title, description = '', priority = 1) {
        const task = {
            id: this.nextId++,
            title,
            description,
            priority,
            completed: false,
            createdAt: new Date()
        };

        this.tasks.set(task.id, task);
        console.log(`✅ Task added: "${title}" (ID: ${task.id})`);
        return task;
    }

    // Complete a task
    completeTask(id) {
        const task = this.tasks.get(id);
        if (!task) {
            throw new Error(`Task with ID ${id} not found`);
        }

        task.completed = true;
        task.completedAt = new Date();
        console.log(`🎉 Task completed: "${task.title}"`);
        return task;
    }

    // Get tasks by status
    getTasksByStatus(completed = false) {
        return Array.from(this.tasks.values())
            .filter(task => task.completed === completed);
    }

    // Get tasks sorted by priority
    getTasksByPriority() {
        return Array.from(this.tasks.values())
            .sort((a, b) => b.priority - a.priority);
    }

    // Get summary statistics
    getSummary() {
        const tasks = Array.from(this.tasks.values());
        const completed = tasks.filter(t => t.completed).length;
        const pending = tasks.length - completed;
        
        return {
            total: tasks.length,
            completed,
            pending,
            completionRate: tasks.length > 0 ? (completed / tasks.length * 100).toFixed(1) : 0
        };
    }

    // Display all tasks
    displayTasks() {
        console.log(`\n=== ${this.name} - Task List ===`);
        
        if (this.tasks.size === 0) {
            console.log("No tasks found.");
            return;
        }

        this.tasks.forEach(task => {
            const status = task.completed ? '✅' : '⭕';
            const priority = '⭐'.repeat(task.priority);
            console.log(`${status} [${task.id}] ${task.title} ${priority}`);
        });
    }
}

// Demo function
function main() {
    console.log("=== Task Manager Demo ===");

    // Create task manager
    const manager = new TaskManager("My Tasks");

    // Add sample tasks
    manager.addTask("Review code", "Check PR #123", 3);
    manager.addTask("Buy groceries", "Milk, eggs, bread", 2);
    manager.addTask("Fix bug", "Database timeout", 5);

    // Complete a task
    manager.completeTask(1);

    // Display all tasks
    manager.displayTasks();

    // Show summary
    const summary = manager.getSummary();
    console.log(`\nSummary: ${summary.completed}/${summary.total} completed (${summary.completionRate}%)`);

    console.log("Demo completed!");
}

// Run the demo
main();
