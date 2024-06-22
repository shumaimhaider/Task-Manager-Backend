import { Outlet, Link } from "react-router-dom";

const Layout = () => {
  return (
    <>
      {/* <nav className="px-10 ">
        <ul className="bg-purple-500">
          <li className="mx-2 font-bold">
            <Link to="/todostype">View To Do Types</Link>
          </li>
          <li className="mx-2 font-bold">
            <Link to="/createtodo">Create Type</Link>
          </li>
          <li className="mx-2 font-bold">
            <Link to="/createtodotask">Create Task</Link>
          </li>
          <li className="mx-2 font-bold">
            <Link to="/viewtodotask">View To Do Task</Link>
          </li>
        </ul>
      </nav> */}

<nav class="px-2 bg-white border-gray-200 dark:bg-gray-900 dark:border-gray-700">
  <div class="container flex flex-wrap items-center justify-between mx-auto">
    
    <div >
      <ul class="flex flex-col p-4 mt-4 border border-gray-100 rounded-lg bg-gray-50 md:flex-row md:space-x-8 md:mt-0 md:text-sm md:font-medium md:border-0 md:bg-white dark:bg-gray-800 md:dark:bg-gray-900 dark:border-gray-700">
        <li>
          <a href="/createtodo" class="block py-2 pl-3 pr-4 text-gray-700 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-gray-400 md:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent" aria-current="page">Create Task Type</a>
        </li>
      
        <li>
          <a href="/todostype" class="block py-2 pl-3 pr-4 text-gray-700 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-gray-400 md:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent">View Task Type</a>
        </li>
        <li>
          <a href="createtodotask" class="block py-2 pl-3 pr-4 text-gray-700 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-gray-400 md:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent">Create Task</a>
        </li>
        <li>
          <a href="viewtodotask" class="block py-2 pl-3 pr-4 text-gray-700 rounded hover:bg-gray-100 md:hover:bg-transparent md:border-0 md:hover:text-blue-700 md:p-0 dark:text-gray-400 md:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent">View Tasks</a>
        </li>
      </ul>
    </div>
  </div>
</nav>

      <Outlet />
    </>
  )
};
export default Layout;