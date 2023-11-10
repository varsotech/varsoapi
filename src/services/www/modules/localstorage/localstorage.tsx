import { useState } from "react"

function useLocalStorage<T>(key: string, initialValue: T) {
  const [state, setState] = useState<T>((): T => {
    // Initialize the state
    try {
      const value: string = window.localStorage.getItem(key) || ""
      // Check if the local storage already has any values,
      // otherwise initialize it with the passed initialValue
      return value ? (JSON.parse(value) as T) : initialValue
    } catch (error) {
      console.log(error)
    }

    return initialValue
  })

  function setValue(value: T) {
    try {
      // If the passed value is a callback function,
      //  then call it with the existing state.
      window.localStorage.setItem(key, JSON.stringify(value))
      setState(value)
    } catch (error) {
      console.log(error)
    }
  }

  return [state, setValue]
}

export default useLocalStorage
