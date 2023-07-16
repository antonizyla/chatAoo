import type { Actions } from "@sveltejs/kit";

export const actions = {

    createUser: async ({ request }) => {
        // will create a new chat in database and return uuid
        const data = await request.formData();
        const name = data.get("name");

        // create a request to api to get new chat uuid
        let user = await fetch(`http://localhost:8081/createUser?username=${name}`).then(res => res.json());

        return { user }
    }

} satisfies Actions;
