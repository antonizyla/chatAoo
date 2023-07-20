import type { PageLoad } from "./$types"

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

}) satisfies PageLoad;

