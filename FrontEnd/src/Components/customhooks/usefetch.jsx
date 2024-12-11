import { useState, useEffect } from "react";

const useGetData = (url) => {
  const [todoItems, setToDoItems] = useState(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    fetch(url)
      .then((res) => res.json())
      .then((res) => {
        console.log(res);
        setToDoItems(res);
        setIsLoading(false);
      })
      .catch((err) => {
        console.error(err);
        setIsLoading(false);
      });
  }, [url]);

  return { todoItems, isLoading };
};

export default useGetData;
