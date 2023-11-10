import { setAxiosToken } from "api/axios"
import Layout from "components/Layout/Layout"
import { useRouter } from "next/router"
import React, { createContext, useContext, useEffect, useReducer } from "react"

type AuthState = {
  token: string | undefined
}

const initialState: AuthState = {
  token: undefined, // undefined means it's not loaded. empty string means it's missing.
}

const AuthContext = createContext<{ state: AuthState; dispatch: Dispatch } | undefined>(undefined)

type Action = {
  type: "setToken"
  token: string
}
type Dispatch = (action: Action) => void

function authReducer(state: AuthState, action: Action): AuthState {
  switch (action.type) {
    case "setToken": {
      return { ...state, token: action.token }
    }
    default: {
      throw new Error(`Unhandled action: ${action}`)
    }
  }
}

type AuthProviderProps = {
  children: React.ReactNode
}

export function AuthProvider({ children }: AuthProviderProps) {
  const [state, dispatch] = useReducer(authReducer, initialState)
  const value = { state, dispatch }

  useEffect(() => {
    const fetchToken = async () => {
      const fetchedToken = localStorage.getItem("token")
      dispatch({ type: "setToken", token: fetchedToken || "" })
    }
    fetchToken()
  }, [])

  return (
    <AuthContext.Provider value={value}>
      <AuthRedirection>{children}</AuthRedirection>
    </AuthContext.Provider>
  )
}

type AuthRedirectionProps = {
  children: React.ReactNode
}

function AuthRedirection({ children }: AuthRedirectionProps) {
  // const [token, setToken] = useToken()
  // const router = useRouter()

  // useEffect(() => {
  //   if (token) {
  //     setAxiosToken(token, () => {
  //       setToken("")
  //     })
  //   }
  // }, [token, setToken])

  // if (token === undefined) {
  //   return <div></div>
  // }

  // if (token === "" && router.pathname !== "/login") {
  //   router.push("/login")
  //   return <div></div>
  // }

  return children
}

function useAuth(): [AuthState, Dispatch] {
  const context = useContext(AuthContext)
  if (context === undefined) {
    throw new Error("useAuth must be used within a AuthContext")
  }
  return [context.state, context.dispatch]
}

export function useToken(): [string | undefined, (token: string) => void] {
  const [state, dispatch] = useAuth()

  function actionFunction(token: string) {
    localStorage.setItem("token", token)
    dispatch({ type: "setToken", token: token })
  }

  return [state.token, actionFunction]
}

export default AuthContext
