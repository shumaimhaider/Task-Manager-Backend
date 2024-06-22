
import React, { useState, useEffect } from "react";
import { Button, TextInput, Checkbox, Label, Modal } from "flowbite-react";
import axios from "axios";
import LoadingSpinner from "../spinner/spinner";

const Task = (item) => {
    console.log(item)
    const [detail, setDetails] = useState("");
    const [isCompleted, setisCompleted] = useState();
    const [rselected,setRselected]=useState(item.is_completed)
    const [todoItems, setToDoItems] = useState([])
    const [options, setOptions] = useState("")
    const [open, setOpen] = useState(false)
    const[isLoading,setisLoading]=useState(false)


    const getData = () => {
        fetch('http://localhost:8080/Get/todos/Types')
            .then((res) => res.json())
            .then((res) => {
                console.log(res)
                setToDoItems(res)
                console.log("r selected",rselected)
            })
    };
    useEffect(() => {
        getData()
    }, []);

    const items = todoItems.map((todoItem, i) => <option key={i} selected={item.id === todoItem.id}>{todoItem.todo_type_name}</option>)
    const handleSubmit = async (e) => {
        setisLoading(true)
        e.preventDefault();
        try {
            const data = { todo_task_id: item.todo_task_id, details: detail,todo_type_name:options,is_completed: isCompleted === "true" }
            let res = await axios.patch("http://localhost:8080/update/todo/task/completed", data)
            console.log(res.data)
            if (res.status == 200) {
                setDetails("");
                setisCompleted("false")
                setOptions("")
                setTimeout(() => {
                    setisLoading(false)
                }, 1000);
                alert("todo updated successfully")
            } else {
                alert("Error Updating!!!")
            }
        }catch (e) {
            alert(e)
        }
    };
    const handleOpen = () => {
        setOpen(true)
    }
    const handleClose = () => {
        setOpen(false)
    }
    return (
        <React.Fragment>
            <Button onClick={handleOpen}>
                Update
            </Button>
            <Modal
                show={open}
                size="md"
                popup={true}
                onClose={handleClose}
            >
                <Modal.Header />
                <Modal.Body>
                {isLoading && <LoadingSpinner/>}

                    <div className="space-y-6 px-6 pb-4 sm:pb-6 lg:px-8 xl:pb-8">
                        <h3 className="text-xl font-medium text-gray-900 dark:text-white">
                            Edit your Tasks
                        </h3>
                        <div>
                            <div className="mb-2 block">
                                <Label
                                    htmlFor="Type"
                                    value="Details"
                                />
                            </div>
                            <TextInput
                                id="type"
                                placeholder="type"
                                required={true}
                                defaultValue={item.details}
                                onChange={(e) => setDetails(e.target.value)}
                            />
                            <div className="mb-2 block">
                                <Label
                                    htmlFor="Type"
                                    value="Select Type"
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
                            {/* checked={rselected == true} */}
                            <div onChange={(e) => setisCompleted(e.target.value)} >
                                <input type="radio" value="true" name="task"  /> <label >Completed</label>
                                <input type="radio" value="false" name="task"  /><label >Not Completed</label>
                            </div>
                        </div>
                        <br />
                        <Button color="success" onClick={handleSubmit}>
                            Update
                        </Button>
                    </div>
                </Modal.Body>
            </Modal>
        </React.Fragment>
    )
}
export default Task
