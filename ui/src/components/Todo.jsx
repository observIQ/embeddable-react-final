import React, { useState } from "react";
import { TrashCanIcon } from "./TrashCanSvg";

export const Todo = (props) => {
  const [completed, setCompleted] = useState(props.todo.completed);
  const [showDelete, setShowDelete] = useState(false);

  async function handleCheckClick(e) {
    const payload = {
      completed: e.target.checked,
    };

    const resp = await fetch(`/api/todos/${props.todo.id}`, {
      method: "PUT",
      body: JSON.stringify(payload),
    });

    const body = await resp.json();
    const { todo } = body;
    setCompleted(todo.completed);
  }

  async function handleDelete(e) {
    await fetch(`/api/todos/${props.todo.id}`, { method: "DELETE" });
    props.onDeleteSuccess();
  }
  return (
    <div
      className="todo"
      onMouseEnter={() => setShowDelete(true)}
      onMouseLeave={() => setShowDelete(false)}
    >
      <input
        type={"checkbox"}
        className="checkbox"
        checked={completed}
        onChange={handleCheckClick}
      />
      <p>{props.todo.description}</p>

      {showDelete && (
        <button className="delete" onClick={handleDelete}>
          <TrashCanIcon />
        </button>
      )}
    </div>
  );
};
