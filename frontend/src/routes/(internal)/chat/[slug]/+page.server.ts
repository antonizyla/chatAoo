import type { PageServerLoad } from './$types';

export const load: PageServerLoad = (async ({ params }) => {
    const chatId = params.slug;
    // check with api if chat exists
    // if chat exists, connect to chat otherwise give link to create one
    const chat = await fetch(`http://localhost:8081/chats/${chatId}`)

    if (chat.status != 400) {
        const chatData = await chat.json();
        return { chat: chatData, exists: true };
    } else {
        return { chat: null, exists: false };
    }

}) 
