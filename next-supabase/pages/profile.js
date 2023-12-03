// pages/profile.js
import { Typography, Button } from "@supabase/ui";
const { Text } = Typography
import { Auth } from '@supabase/auth-ui-react';
import { ThemeSupa } from '@supabase/auth-ui-shared'
import { createClientComponentClient } from '@supabase/auth-helpers-nextjs'
import { supabase } from '../api'

function Profile(props) {
    const { user } = Auth.useUser();
    if (user)
        return (
            <>
                <Text>Signed in: {user.email}</Text>
                <Button block onClick={() => props.supabaseClient.auth.signOut()}>
                    Sign out
                </Button>
            </>
        );
    return props.children
}

export default function AuthProfile() {
    return (
        <Auth.UserContextProvider supabaseClient={supabase}>
            <Profile supabaseClient={supabase}>
                <Auth supabaseClient={supabase} />
            </Profile>
        </Auth.UserContextProvider>
    )
}
