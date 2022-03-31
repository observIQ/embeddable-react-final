import React, { useState } from "react";

export const NewTodoInput = (props) => {
  const [showInput, setShowInput] = useState(false);
  const [value, setValue] = useState("");

  async function handleSave(e) {
    e.preventDefault();

    const payload = {
      description: value,
    };

    const resp = await fetch("/api/todos", {
      method: "POST",
      body: JSON.stringify(payload),
    });

    const body = await resp.json();
    console.log({ body });

    props.onCreateSuccess(body.todo);
    setShowInput(false);
    setValue("");
  }

  return showInput ? (
    <form className="input-box">
      <input
        className="todo-input"
        autoFocus
        value={value}
        onChange={(e) => setValue(e.target.value)}
        type="text"
        placeholder="To do..."
      />
      <button className="save-button" onClick={(e) => handleSave(e)}>
        Save
      </button>
    </form>
  ) : (
    <div className="button-box">
      <button className="new-button" onClick={() => setShowInput(true)}>
        New
      </button>
    </div>
  );
};
