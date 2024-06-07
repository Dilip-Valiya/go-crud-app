import React from "react";

interface Todo {
  id: number;
  title: string;
  description: string;
}

async function getData() {
  const res = await fetch("http://localhost:8000/todos", { cache: "no-store" });
  const todos = await res.json();
  return todos;
}

export default async function Todos() {
  const todos = await getData();

  return (
    <div className="container mx-auto p-4 h-full">
      <h1 className="text-3xl font-bold mb-4">Todos</h1>
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
        {todos.map((todo: Todo) => (
          <div key={todo.id} className="p-4 border rounded-lg shadow-md">
            <h2 className="text-xl font-bold">{todo.title}</h2>
            <p className="text-gray-600">{todo.description}</p>
          </div>
        ))}
      </div>
    </div>
  );
}
