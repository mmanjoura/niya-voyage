const Social = () => {
  const socialContent = [
    { id: 1, icon: "icon-facebook", link: "https://www.facebook.com/niyavoyage" },
    { id: 2, icon: "icon-twitter", link: "https://twitter.com/niyavoyage" },
    { id: 3, icon: "icon-instagram", link: "https://www.instagram.com/niyavoyage/" },
    { id: 5, icon: "icon-tumblr", link: "https://www.tumblr.com/niyavoyage" },
    { id: 6, icon: "icon-telegram", link: "https://t.me/niyavoyage" },
    { id: 7, icon: "icon-tiktok", link: "https://www.tiktok.com/@niyavoyage" },
    { id: 8, icon: "icon-youtube", link: "https://www.youtube.com/@niyavoyage" },
    { id: 9, icon: "icon-pinterest", link: "https://www.pinterest.com/niyavoyage/" },
     { id: 10, icon: "icon-snapchat", link: "https://www.snapchat.com/add/niyavoyage" },
     { id: 11, icon: "icon-threads", link: "https://www.threads.net/@niyavoyage" }
  ];
  return (
    <>
      {socialContent.map((item) => (
        <a
          href={item.link}
          target="_blank"
          rel="noopener noreferrer"
          key={item.id}
        >
          <i className={`${item.icon} text-14`} />
        </a>
      ))}
    </>
  );
};

export default Social;
