import React from "react"
import { Label, Checkbox,TextInput,Button,Card } from "flowbite-react"
import { useState } from "react";
import axios from 'axios'
import LoadingSpinner from "../spinner/spinner";
function FormCreateType() {
  const [toDoType, setToDoType] = useState("");
  const [message, setMessage] = useState("");
  const [isLoading,setisLoading]=useState(false)

  let handleSubmit = async (e) => {
    setisLoading(true)
    e.preventDefault();
    try {
      const data = { todo_type_name: toDoType }
      let res = await axios.post("http://localhost:8080/create/todo/type", data)
      console.log(res.data)
      if (res.status == 200) {
        setToDoType("");
        setTimeout(() => {
          setisLoading(false)
        }, 1000);
        setMessage("To do type created successfully");
      } else {
        setMessage("Some error occured");
      }
    } catch (e) {
      alert(e)
    }
  };

  return (  
 <div >
  <Card className=" flex justify-center">
  {isLoading && <LoadingSpinner />}
    <form className="flex flex-col gap-4" onSubmit={handleSubmit}>
      <div>
        <div className="mb-2 block">
          <Label
            htmlFor="task type"
            value="Enter task type "
          />
        </div>
        <TextInput
          id="task"
          type="task"
          placeholder="Enter your task type"
          required={true}
          onChange={(e) => setToDoType(e.target.value)}
        />
      </div>
      <Button type="submit">
        Submit
      </Button>
      <div>{message ? <p>{message}</p> : null}</div>
    </form>
  </Card>
</div>
    )
}

export default FormCreateType