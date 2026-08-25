

export type TaskStatus = "pending" | "done";



export type Task = {
    id: string;
    title: string;
    status: TaskStatus;
    createdAt: Date;
    updatedAt: Date;
};

export type CreateTaskInput = {
    title: string;

};


export type UpdateTaskInput = {
    title?: string;
    status: TaskStatus;
};


export type PublicTask = Omit<Task, "createdAt" | "updatedAt"> &{
    createdAt: string;
    updatedAt: string;
};


