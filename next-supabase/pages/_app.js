// pages/_app.js
import Link from 'next/link'
import { useState, useEffect } from 'react'
import { supabase } from '../api'
import '../styles/globals.css'

function MyApp({ Component, pageProps }) {
  const [user, setUser] = useState(null);
  useEffect(() => {
    // const { data: authListener } = supabase.auth.onAuthStateChange(
    const { data: { subscription } } = supabase.auth.onAuthStateChange(
      async () => checkUser()
    )
    checkUser()

    return () => {
      // console.log("authListener: ", authListener);

      // authListener?.unsubscribe()
      subscription?.unsubscribe();
    };
  }, [])
  async function checkUser() {
    const { data: { user } } = await supabase.auth.getUser();
    console.log("got user: ", user);
    setUser(user)
  }
  return (
    <div>
      <nav className="p-6 border-b border-gray-300">
        <Link href="/">
          <span className="mr-6 cursor-pointer">Home</span>
        </Link>
        {
          user && (
            <Link href="/create-post">
              <span className="mr-6 cursor-pointer">Create Post</span>
            </Link>
          )
        }
        <Link href="/profile">
          <span className="mr-6 cursor-pointer">Profile</span>
        </Link>
      </nav>
      <div className="py-8 px-16">
        <Component {...pageProps} />
      </div>
    </div>
  )
}

export default MyApp
