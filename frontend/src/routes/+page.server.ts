import type { Actions } from "@sveltejs/kit";

export const actions = {

    createChat: async ({ request }) => {
        // will create a new chat in database and return uuid
        const data = await request.formData();
        const name = data.get("name");
        const description = data.get("description");

        // create a request to api to get new chat uuid
        let chat = await fetch(`http://localhost:8081/createChat?name=${name}&description=${description}`).then(res => res.json());

        return { uuid: chat.id, name: chat.name, description: chat.description, success: true };

    }

} satisfies Actions;
