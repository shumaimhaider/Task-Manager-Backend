import React from "react";
import { Label, Dropdown, TextInput, Button } from "flowbite-react"
import { useState, useEffect } from "react";
import axios from 'axios'
import LoadingSpinner from "../spinner/spinner";

function CreateTask() {
  const [detail, setDetails] = useState("");
  const [isCompleted, setisCompleted] = useState("false");
  // cons[toDoType,setToDoType]=(5)
  const [todoItems, setToDoItems] = useState([])
  const [options, setOptions] = useState("")
  const [message, setMessage] = useState("");
  const[isLoading,setisLoading]=useState(false)
  ///////////////////////////////////////////////
  useEffect(() => {
   
    fetch('http://localhost:8080/Get/todos/Types')
      .then((res) => res.json())
      .then((res) => {
        console.log(res)
        setToDoItems(res)
     
      })
  }, []);
  let items = todoItems.map((item, i) => <option key={i}>{item.todo_type_name}</option>)
  ////////////////////////////////////////////////////////////////////////////////////////////////////
  let handleSubmit = async (e) => {
    setisLoading(true)
    e.preventDefault();
    try {
      console.log(isCompleted)
      debugger
      const data = { details: detail, is_completed: isCompleted === "true", todo_type_name: options }
      console.log(isCompleted)
      let res = await axios.post("http://localhost:8080/create/todo/task/name", data)
      console.log(res.data)
      if (res.status == 200){
        alert("todo created successfully")
        setDetails("");
        setisCompleted("false")
        setOptions("")
        setTimeout(() => {
          setisLoading(false)
        }, 1000);
        setMessage("To Do task created successfully");
      } else {
        alert("Error creating!!!")
        setMessage("Some error occured");
      }
    } catch (e) {
      alert(e)
    }
  };
  return (
    <>
      {isLoading && <LoadingSpinner />}

    
    <form className="flex flex-col gap-4" onSubmit={handleSubmit}>
      <br />
      <br />
      <div>
        <div className="mb-2 block">
          <Label
            htmlFor="task"
            value="Enter Task detail"
          />
        </div>
        <TextInput
          id="m"
          type="task"
          placeholder="Task Detail"
          required={true}
          onChange={(e) => setDetails(e.target.value)}
                                                                                                        
        />
      </div>

      <select onChange={(e) => setOptions(e.target.value)}>
          {items}
        </select>

      <div className="mb-2 block">
        <Label
          htmlFor="status"
          value="Task Status"
        />
      </div>
      <div onChange={(e) => setisCompleted(e.target.value)}  >
        <input type="radio" value="true" name="task" /> <label className="">Completed</label>
        <input type="radio" value="false" name="task" /><label className="">Not Completed</label>
      </div>
      <Button type="submit">
        Submit
      </Button>
    </form>
    </>
  )
}
export default CreateTask