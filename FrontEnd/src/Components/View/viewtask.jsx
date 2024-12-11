import React from "react";
import { useState, useEffect } from "react";
import '../../App.css'
import { Table, Button } from 'flowbite-react'
import axios from "axios";
import Task from "../modals/task";
import LoadingSpinner from "../spinner/spinner";
import useGetData from "../customhooks/usefetch";

const Todostask = () => {
  const { isLoading, todoItems } = useGetData('http://localhost:8080/Get/todos/Tasks')

  return (
    <>
      {isLoading ? <LoadingSpinner /> :
        <Table>
          <Table.Head>
            <Table.HeadCell>
              task details
            </Table.HeadCell>
            <Table.HeadCell>
              Type
            </Table.HeadCell>
            <Table.HeadCell>
              Status
            </Table.HeadCell>
            <Table.HeadCell>
              Actions
            </Table.HeadCell>
          </Table.Head>
          <Table.Body className="divide-y">
            {todoItems?.map((item, i) => {
              return (
                <Table.Row className="bg-white dark:border-gray-700 dark:bg-gray-800">
                  <Table.Cell className="whitespace-nowrap font-medium text-gray-900 dark:text-white">
                    {item.details}
                  </Table.Cell>
                  <Table.Cell>
                    {item.CreateToDoType.todo_type_name}
                  </Table.Cell>
                  <Table.Cell>
                    <div>{item.is_completed ? <p>completed</p> : <p>Not Completed</p>}</div>
                  </Table.Cell>
                  <Table.Cell className="text-gray-900">
                    {/* <button onClick={() => deleteSubmit(item)} className="bg-yellow-400 "> delete </button> */}
                    {/* <Button color="success" onClick={() => deleteSubmit(item) } disabled={isLoading}>Delete</Button> */}
                    <br />
                    <Task {...item} />
                  </Table.Cell>
                </Table.Row>
              )
            })}
          </Table.Body>
        </Table>}

    </>
  )

}
export default Todostask