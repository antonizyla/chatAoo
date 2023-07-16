<script lang="ts">
	import type { ActionData, PageData } from './$types';
	export let data: PageData;
	export let form: ActionData;

</script>

<h1>Welcome to Chat App</h1>
{#if !form?.success}
	<form action="?/createChat" method="POST">
		<input type="text" name="name" placeholder="Chat Name" />
		<input type="text" name="description" placeholder="Chat Description" />
		<button type="submit"> Create Chat </button>
	</form>
{:else}
	Successfully Created Chat with name {form?.name} and url: /chat/{form.uuid}
	<button><a href={`/chat/${form.uuid}`}>Launch Chat</a></button>
{/if}
