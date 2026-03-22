<script lang="ts">
	import type { PageProps } from './$types';
	import { getLobbyQuery } from '$api/lobby';

	let { params }: PageProps = $props();

	const lobbyQuery = getLobbyQuery(() => params.id);
</script>


{#if lobbyQuery.isLoading}
Loading...
{/if}

{#if lobbyQuery.isError}
	{#if lobbyQuery.error.status === 404}
		<p>Lobby not found.</p>
	{:else}
		<p>An error occurred while fetching the lobby.</p>
	{/if}
{/if}

{#if lobbyQuery.isSuccess}
	<div class="mx-auto flex max-w-1/2 flex-col gap-4 mt-10">
			<h1 class="text-4xl text-center">{lobbyQuery.data.name}</h1>
			<p>Game Edition: {lobbyQuery.data.gameEditionId}</p>
			<p>WIP!</p>

		<ul>
			TODO
			<li>- Add encounters</li>
			<li>- Add players</li>
			<li>- Add utility tools (optional)</li>
			<li>- Some nice UI/UX</li>
		</ul>
	</div>
{/if}
