import React from "react";
import { useState, useEffect } from "react";
import { Button, Table } from 'flowbite-react'
import axios from "axios";
import Type from "../modals/type";
import LoadingSpinner from "../spinner/spinner";
import usegetData from "../customhooks/usefetch";
const Todos = () => {
  const [toDoItems, setToDoItems] = useState([])
  // const [isLoading, setIsLoading] = useState(true);
  const { isLoading, todoItems } = usegetData('http://localhost:8080/Get/todos/Types');
  
  const getData = () =>{
    fetch('http://localhost:8080/Get/todos/Types')
      .then((res) => res.json())
      .then((res) => {
        setToDoItems(res)
        // setIsLoading(false)
      }).catch(err => {
        // setIsLoading(false)
      })
  }
  useEffect(() => {
    setTimeout(() => {
      getData()
    }, 2000);
  }, [todoItems]);

  let deleteSubmit = async (item) => {
    // setIsLoading(true)
    try {
      console.log(item.todo_type_id)
      // const data = {item.todo_type_id}
      const data = { "todo_type_id": item.todo_type_id }
      debugger
      let res = await axios.post("http://localhost:8080/delete/todo/type", data)
      // console.log(res.data)
      if (res.status == 200) {
        alert("Type deleted Sucessfully")
        setTimeout(() => {
          // setIsLoading(false)
      }, 1000);
        // getData()
      } else {
        alert("Error while deleting!!!")
      }
    } catch (e) {
      alert(e)
    }
  };
  return (
    <>
      {isLoading && <LoadingSpinner/>}
      <Table>
        <Table.Head>
          <Table.HeadCell>
            To Types Names
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
                  {item.todo_type_name}
                </Table.Cell>
                <Table.Cell className="text-gray-900">
                  <Button color="success" onClick={() => deleteSubmit(item)}>
                    Delete
                  </Button>
                  <br />
                  <Type {...item} />
                </Table.Cell>
              </Table.Row>
            )
          })}
        </Table.Body>
      </Table>

    </>
  )
}
export default Todos