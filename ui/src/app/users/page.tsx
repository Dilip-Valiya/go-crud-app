import UserCard from "@/components/UserCard";
import React from "react";

interface User {
  id: number;
  name: string;
  email: string;
}

export default async function Users() {
  const users = await getData();

  return (
    <div className="container mx-auto p-4 h-full">
      <h1 className="text-3xl font-bold mb-4">Users</h1>
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
        {users.map((user: User) => (
          <UserCard key={user.id} name={user.name} email={user.email} />
        ))}
      </div>
    </div>
  );
}

async function getData() {
  const res = await fetch("http://localhost:8000/users");
  const users = await res.json();
  return users;
}
