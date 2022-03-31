import React, { useState } from "react";
import { useCallback } from "react";
import { useEffect } from "react";
import { NewTodoInput } from "./NewTodoForm";
import { Todo } from "./Todo";
import { Link } from "react-router-dom";

export const Todos = () => {
  const [todos, setTodos] = useState([]);

  const fetchTodos = useCallback(async () => {
    const resp = await fetch("/api/todos");
    const body = await resp.json();
    const { todos } = body;

    setTodos(todos);
  }, [setTodos]);

  useEffect(() => {
    fetchTodos();
  }, [fetchTodos]);

  function onDeleteSuccess() {
    fetchTodos();
  }

  function onCreateSuccess(newTodo) {
    setTodos([...todos, newTodo]);
  }

  return (
    <>
      <h3>To Do:</h3>
      <div className="todos">
        {todos.map((todo) => (
          <Todo key={todo.id} todo={todo} onDeleteSuccess={onDeleteSuccess} />
        ))}
      </div>
      <NewTodoInput onCreateSuccess={onCreateSuccess} />
      <Link to="/about" className="nav-link">
        Learn more...
      </Link>
    </>
  );
};
