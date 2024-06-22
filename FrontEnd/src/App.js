import logo from './logo.svg';

import { BrowserRouter,Route,Routes } from 'react-router-dom';
import FormCreateType from './Components/create/createtype';
import Layout from './Components/layout/layout'
import CreateTask from './Components/create/createtask';
import ViewToDo from './Components/View/viewtype';
import ViewTask from './Components/View/viewtask';
function App() {
  return (
    <div>
    <BrowserRouter>
  
     <Layout />
     
      <Routes>
        <Route path="/" element={<FormCreateType />} />
        <Route path="todostype" element={<ViewToDo />} />
        <Route path="createtodo" element={<FormCreateType />} />
        <Route path="createtodotask" element={<CreateTask/>} />
        <Route path="viewtodotask" element={<ViewTask/>} />
      </Routes>
    </BrowserRouter>
    </div>
  )
}

export default App;
