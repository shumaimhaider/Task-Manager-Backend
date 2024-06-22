import React, { useState } from "react";
import { Button, TextInput, Checkbox, Label, Modal } from "flowbite-react";
import axios from "axios";
import LoadingSpinner from "../spinner/spinner";
const Type = (items) => {
    console.log(items.todo_type_name)

    const [open, setOpen] = useState(false)
    const[isLoading,setisLoading]=useState(false)

    const handleOpen = () => {
        setOpen(true)
    }

    const handleClose = () => {
        setOpen(false)
    }
    ////////////////////////////////////////////
    const [toDoType, setToDoType] = useState("");
    
    let handleSubmit = async (e) => {
        setisLoading(true)
        e.preventDefault();
        try {
            const data = { todo_type_id: items.todo_type_id, todo_type_name: toDoType }
            let res = await axios.patch("http://localhost:8080/update/todo/type", data)
            console.log(res.data)
            if (res.status == 200) {
                setTimeout(() => {
                    setisLoading(false)
                }, 1000);
                alert("Updated Sucessfully")
            } else {
                alert("Didn't Updated")
            }
        } catch (e) {
            alert(e)
        }
    };

    return (
        <React.Fragment>
            <Button onClick={handleOpen}>
                Edit
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
                            Edit the Type
                        </h3>
                        <div>
                            <div className="mb-2 block">
                                <Label
                                    htmlFor="Type"
                                    value="Type"
                                />
                            </div>
                            <TextInput
                                id="type"
                                placeholder="type"
                                required={true}
                                defaultValue={items.todo_type_name}
                                // value={items.todo_type_name}
                                onChange={(e) => setToDoType(e.target.value)}
                            />
                        </div>
                        <br />
                        <div className="w-full">
                            <Button onClick={handleSubmit}>
                                Update
                            </Button>
                        </div>
                    </div>
                </Modal.Body>
            </Modal>
        </React.Fragment>
    )
}
export default Type