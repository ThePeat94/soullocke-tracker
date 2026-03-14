import { defineConfig } from '@hey-api/openapi-ts';

export default defineConfig({
	input: '../openapi/openapi.yaml',
	output: 'src/api/generated',
	plugins: [
		'@hey-api/client-fetch',
		{
			name: '@tanstack/svelte-query',
			queryOptions: true,
			mutationOptions: true,
			queryKeys: true
		}
	]
});
