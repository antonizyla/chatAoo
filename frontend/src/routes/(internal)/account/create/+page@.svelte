<script lang="ts">
	import { enhance } from '$app/forms';
	import Button from '$lib/svelte-components/components/Button/Button.svelte';
	import { onMount } from 'svelte';
	import type { ActionData } from './$types';
	import { previousPage } from '$lib/stores/previousPageStore';

	export let form: ActionData;

	let mounted: boolean = false;
	let username: string | null = '';
	onMount(async () => {
		mounted = true;
		username = localStorage.getItem('userName');
		if (username) {
			window.location.href = '/account';
		}
	});

	async function linkUser(chatID: string, userID: string) {
		console.log('linking user');
		const link = await fetch(`http://localhost:8081/chats/link`, {
			method: 'POST',
			body: JSON.stringify({ chat_id: chatID, user_id: userID })
		}).then((res) => {
			res.status, res.statusText;
		});
		console.log(link);
		// redirect to chatpage
		window.location.href = `/chat/${chatID}`;
		localStorage.removeItem('chatId');
	}

	$: {
		if (form?.user && mounted) {
			localStorage.setItem('userName', form.user.name);
			localStorage.setItem('userID', form.user.user_id);
			const prevPage = localStorage.getItem('previousPage');
			if (prevPage) {
				window.location.href = prevPage;
				// @ts-ignore, says it's possibly null, really???
				let chatID: string | null = /((\w{4,12}-?)){5}/.exec(prevPage)[0];
				if (chatID) {
					linkUser(chatID, form.user.user_id);
					localStorage.removeItem('previousPage');
				}
			}
			username = form.user.username;
		}
	}
</script>

<div class="text-lg mx-auto max-w-prose text-center w-fit p-8">
	No User Identifier has been found in your browser, Please Create one below
</div>
<div class="flex flex-col gap-2 pt-8 mx-auto w-fit">
	<p class="py-2 text-center">Enter a User Identifier to Create an Account</p>
	<form action="?/createUser" method="POST" use:enhance class="flex flex-row gap-2 items-center">
		<input
			class="p-1.5 border border-text border-solid rounded-md"
			type="text"
			name="name"
			id="user"
			placeholder="Enter a user identifier"
			autofocus
			required
		/>
		<button type="submit" on:click={() => {}}
			><Button primary size="small">Create Account</Button></button
		>
	</form>
</div>
