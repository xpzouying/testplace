// pages/create-post.js
import { useState } from 'react'
import { v4 as uuid } from 'uuid'
import { useRouter } from 'next/router'
import dynamic from 'next/dynamic'
import "easymde/dist/easymde.min.css"
import { supabase } from '../api'

const SimpleMDE = dynamic(() => import('react-simplemde-editor'), { ssr: false })
const initialState = { title: '', content: '' }

function CreatePost() {
    const [post, setPost] = useState(initialState)
    const { title, content } = post
    const router = useRouter()
    function onChange(e) {
        setPost(() => ({ ...post, [e.target.name]: e.target.value }))
    }
    async function createNewPost() {

        try {

            if (!title || !content) return
            const { data: { user } } = await supabase.auth.getUser();

            console.log('user: ', user);

            const id = uuid()
            post.id = id
            const { data, error } = await supabase
                .from('posts')
                .insert([
                    { title, content, user_id: user.id, user_email: user.email }
                ])
                .single()

            console.log('result: ', result);

            if (error) throw error;

            console.log('data: ', data);

            // NOTE(zy): 这里不生效
            router.push(`/posts/${data.id}`)
        } catch (error) {
            console.error('error catch: ', error);
        }
    }
    return (
        <div>
            <h1 className="text-3xl font-semibold tracking-wide mt-6">Create new post</h1>
            <input
                onChange={onChange}
                name="title"
                placeholder="Title"
                value={post.title}
                className="border-b pb-2 text-lg my-4 focus:outline-none w-full font-light text-gray-500 placeholder-gray-500 y-2"
            />
            <SimpleMDE
                value={post.content}
                onChange={value => setPost({ ...post, content: value })}
            />
            <button
                type="button"
                className="mb-4 bg-green-600 text-white font-semibold px-8 py-2 rounded-lg"
                onClick={createNewPost}
            >Create Post</button>
        </div>
    )
}

export default CreatePost
