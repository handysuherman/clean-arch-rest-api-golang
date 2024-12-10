ALTER USER 'mock_user'@'%' 
    REQUIRE ISSUER '/C=ID/O=xyz-multifinance/L=Denpasar/CN=xyz-multifinance-mysql-ca' 
    AND SUBJECT '/C=ID/O=xyz-multifinance/L=Denpasar/CN=xyz-multifinance-mysql-client';
GRANT ALL PRIVILEGES ON *.* TO 'mock_user'@'%';
FLUSH PRIVILEGES;
