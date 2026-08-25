

export type TaskStatus = "pending" | "done";



export type Task = {
    id: string;
    title: string;
    status: TaskStatus;
    createdAt: Date;
    updatedAt: Date;
};

export type CreateTaskInput = Pick<Task, "title">;


export type UpdateTaskInput = Partial<Pick<Task, "title" | "status">>;


export type PublicTask = Omit<Task, "createdAt" | "updatedAt"> &{
    createdAt: string;
    updatedAt: string;
};


