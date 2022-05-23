import React, { useEffect, useState } from "react";

import axios from "axios";

const ToDoList: React.FC = () => {
    const [task, setTask] = useState<string | number>();
    const [items, setItems] = useState([]);

    const getTask = () => {};
    const onSubmit = () => {};
    const onChange = () => {};

    useEffect(() => {
        getTask();
    }, []);

    return (
        <div>
            <div className="todo_row">
                <div className="todo_header" color="yellow">
                    TO DO LIST
                </div>
            </div>
            <div className="todo_row">
                <form onSubmit={onSubmit}>
                    <input
                        type={"text"}
                        name={"task"}
                        onChange={onChange}
                        value={task}
                        placeholder="Create task"
                    />
                </form>
            </div>
        </div>
    );
};

export default ToDoList;
