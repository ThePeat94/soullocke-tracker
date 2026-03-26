import { createQuery } from '@tanstack/svelte-query';
import { getEditionsOptions } from '$api/generated/@tanstack/svelte-query.gen';

export const getEditionsQuery = ()  => {
	return createQuery(() => {
		return {
			...getEditionsOptions()
		};
	})
};
