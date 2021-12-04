import type { NextPage } from 'next';
import { useEffect, useState } from 'react';
import Head from 'next/head';
import Button from '@mui/material/Button';

import { connect, sendMessage, SocketData } from '../modules/Home/api/socket';
import Header from '../modules/Home/components/Header';
import ChatHistory from '../modules/Home/components/ChatHistory';

const Home: NextPage = () => {
  const [chatHistory, setChatHistory] = useState<SocketData[]>([]);

  useEffect(() => {
    connect(({ data }) => {
      setChatHistory((prev) => [...prev, data]);
    });
  }, []);

  const handleSend = () => {
    sendMessage({ message: 'Hello from Next' });
  };

  return (
    <div>
      <Head>
        <title>Create Next App</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main>
        <Header />
        <ChatHistory chatHistory={chatHistory} />
        <Button variant="contained" onClick={handleSend}>
          Send Message
        </Button>
      </main>
    </div>
  );
};

export default Home;
