import { Typography } from '@mui/material';
import { SocketData } from '../../api/socket';
import Container from './Container';

type Props = {
  chatHistory: SocketData[];
};

const ChatHistory = ({ chatHistory }: Props) => {
  const messages = chatHistory.map((msg, index) => (
    <Typography key={index}>{msg.message}</Typography>
  ));

  return (
    <Container>
      <Typography variant="h5">Chat History</Typography>
      {messages}
    </Container>
  );
};

export default ChatHistory;
