import { client } from './generated/client.gen';
import { PUBLIC_BACKEND_URL} from '$env/static/public';

client.setConfig({
	baseUrl: PUBLIC_BACKEND_URL,
	credentials: 'include',
});
