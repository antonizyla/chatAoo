import type { PageServerLoad } from "./$types"

import type { Actions } from "@sveltejs/kit";

export const load = (async (event) => {
    const chatId = event.params.slug;
    // check with api if chat exists
    // if chat exists, connect to chat otherwise give link to create one 
    const chat = await fetch(`http://localhost:8081/checkChat?id=${chatId}`).then(res => res.json());

    if (chat.name) {
        return { chat: chat, exists: true }
    } else {
        return { chat: null, exists: false }
    }

}) satisfies PageServerLoad;

export const actions: Actions = {

    generateUser: async ({ request }) => {
        const userData = await request.formData()

        const username = userData.get("username");
        const chatId = userData.get("chatID");

        console.log(chatId)

        const createdUser = await fetch(`http://localhost:8081/createUser?username=${username}&chat=${chatId}`).then(res => res.json());

        if (createdUser.id) {
            return { createdUser: createdUser, username: username, exists: true }
        } else {
            return { createdUser: null, username: null, exists: false }
        }

    }

} satisfies Actions;

