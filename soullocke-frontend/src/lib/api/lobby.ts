import { createQuery } from '@tanstack/svelte-query';
import { getLobbyOptions } from './generated/@tanstack/svelte-query.gen.ts';

export function createLobbyQuery(getId: () => string) {
	return createQuery(() => {
		const id = getId();

		return {
			...getLobbyOptions({ path: { id } }),
			enabled: id.length > 0
		};
	});
}
