PGDMP      #                }            nextmedis-test    17.2 (Debian 17.2-1.pgdg120+1)    17.2 (Debian 17.2-1.pgdg120+1) 7    n           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                           false            o           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                           false            p           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                           false            q           1262    17886    nextmedis-test    DATABASE     {   CREATE DATABASE "nextmedis-test" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';
     DROP DATABASE "nextmedis-test";
                     raihan    false                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
                     pg_database_owner    false            r           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                        pg_database_owner    false    4            �            1259    18024    accounts    TABLE     �   CREATE TABLE public.accounts (
    id integer NOT NULL,
    user_id integer NOT NULL,
    balance numeric(15,2) DEFAULT 0 NOT NULL
);
    DROP TABLE public.accounts;
       public         heap r       raihan    false    4            �            1259    18023    accounts_id_seq    SEQUENCE     �   CREATE SEQUENCE public.accounts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.accounts_id_seq;
       public               raihan    false    228    4            s           0    0    accounts_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.accounts_id_seq OWNED BY public.accounts.id;
          public               raihan    false    227            �            1259    18004 
   cart_items    TABLE     g  CREATE TABLE public.cart_items (
    id integer NOT NULL,
    user_id integer NOT NULL,
    product_id integer NOT NULL,
    quantity integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT cart_items_quantity_check CHECK ((quantity > 0))
);
    DROP TABLE public.cart_items;
       public         heap r       raihan    false    4            �            1259    18003    cart_items_id_seq    SEQUENCE     �   CREATE SEQUENCE public.cart_items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.cart_items_id_seq;
       public               raihan    false    226    4            t           0    0    cart_items_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.cart_items_id_seq OWNED BY public.cart_items.id;
          public               raihan    false    225            �            1259    17983    order_details    TABLE     �  CREATE TABLE public.order_details (
    id integer NOT NULL,
    order_id integer NOT NULL,
    product_id integer NOT NULL,
    quantity integer NOT NULL,
    price numeric(10,2) NOT NULL,
    subtotal numeric(10,2) GENERATED ALWAYS AS (((quantity)::numeric * price)) STORED,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT order_details_quantity_check CHECK ((quantity > 0))
);
 !   DROP TABLE public.order_details;
       public         heap r       raihan    false    4            �            1259    17982    order_details_id_seq    SEQUENCE     �   CREATE SEQUENCE public.order_details_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.order_details_id_seq;
       public               raihan    false    224    4            u           0    0    order_details_id_seq    SEQUENCE OWNED BY     M   ALTER SEQUENCE public.order_details_id_seq OWNED BY public.order_details.id;
          public               raihan    false    223            �            1259    17968    orders    TABLE     Z  CREATE TABLE public.orders (
    id integer NOT NULL,
    user_id integer NOT NULL,
    total_price numeric(10,2) NOT NULL,
    status character varying(50) DEFAULT 'pending'::character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);
    DROP TABLE public.orders;
       public         heap r       raihan    false    4            �            1259    17967    orders_id_seq    SEQUENCE     �   CREATE SEQUENCE public.orders_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 $   DROP SEQUENCE public.orders_id_seq;
       public               raihan    false    4    222            v           0    0    orders_id_seq    SEQUENCE OWNED BY     ?   ALTER SEQUENCE public.orders_id_seq OWNED BY public.orders.id;
          public               raihan    false    221            �            1259    17902    products    TABLE     �  CREATE TABLE public.products (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    price numeric(10,2) NOT NULL,
    stock integer NOT NULL,
    category character varying(100),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT products_price_check CHECK ((price >= (0)::numeric)),
    CONSTRAINT products_stock_check CHECK ((stock >= 0))
);
    DROP TABLE public.products;
       public         heap r       raihan    false    4            �            1259    17901    products_id_seq    SEQUENCE     �   CREATE SEQUENCE public.products_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.products_id_seq;
       public               raihan    false    4    220            w           0    0    products_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;
          public               raihan    false    219            �            1259    17889    users    TABLE     �   CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(255) NOT NULL,
    password text NOT NULL,
    email character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);
    DROP TABLE public.users;
       public         heap r       raihan    false    4            �            1259    17888    users_id_seq    SEQUENCE     �   ALTER TABLE public.users ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);
            public               raihan    false    4    218            �           2604    18027    accounts id    DEFAULT     j   ALTER TABLE ONLY public.accounts ALTER COLUMN id SET DEFAULT nextval('public.accounts_id_seq'::regclass);
 :   ALTER TABLE public.accounts ALTER COLUMN id DROP DEFAULT;
       public               raihan    false    227    228    228            �           2604    18007    cart_items id    DEFAULT     n   ALTER TABLE ONLY public.cart_items ALTER COLUMN id SET DEFAULT nextval('public.cart_items_id_seq'::regclass);
 <   ALTER TABLE public.cart_items ALTER COLUMN id DROP DEFAULT;
       public               raihan    false    225    226    226            �           2604    17986    order_details id    DEFAULT     t   ALTER TABLE ONLY public.order_details ALTER COLUMN id SET DEFAULT nextval('public.order_details_id_seq'::regclass);
 ?   ALTER TABLE public.order_details ALTER COLUMN id DROP DEFAULT;
       public               raihan    false    224    223    224            �           2604    17971 	   orders id    DEFAULT     f   ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_id_seq'::regclass);
 8   ALTER TABLE public.orders ALTER COLUMN id DROP DEFAULT;
       public               raihan    false    221    222    222            �           2604    17905    products id    DEFAULT     j   ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);
 :   ALTER TABLE public.products ALTER COLUMN id DROP DEFAULT;
       public               raihan    false    220    219    220            k          0    18024    accounts 
   TABLE DATA                 public               raihan    false    228   �@       i          0    18004 
   cart_items 
   TABLE DATA                 public               raihan    false    226   �@       g          0    17983    order_details 
   TABLE DATA                 public               raihan    false    224   UA       e          0    17968    orders 
   TABLE DATA                 public               raihan    false    222   �A       c          0    17902    products 
   TABLE DATA                 public               raihan    false    220   �E       a          0    17889    users 
   TABLE DATA                 public               raihan    false    218   rI       x           0    0    accounts_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.accounts_id_seq', 1, true);
          public               raihan    false    227            y           0    0    cart_items_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.cart_items_id_seq', 4, true);
          public               raihan    false    225            z           0    0    order_details_id_seq    SEQUENCE SET     B   SELECT pg_catalog.setval('public.order_details_id_seq', 3, true);
          public               raihan    false    223            {           0    0    orders_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.orders_id_seq', 113, true);
          public               raihan    false    221            |           0    0    products_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public.products_id_seq', 15, true);
          public               raihan    false    219            }           0    0    users_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.users_id_seq', 17, true);
          public               raihan    false    217            �           2606    18030    accounts accounts_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.accounts DROP CONSTRAINT accounts_pkey;
       public                 raihan    false    228            �           2606    18032    accounts accounts_user_id_key 
   CONSTRAINT     [   ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_user_id_key UNIQUE (user_id);
 G   ALTER TABLE ONLY public.accounts DROP CONSTRAINT accounts_user_id_key;
       public                 raihan    false    228            �           2606    18012    cart_items cart_items_pkey 
   CONSTRAINT     X   ALTER TABLE ONLY public.cart_items
    ADD CONSTRAINT cart_items_pkey PRIMARY KEY (id);
 D   ALTER TABLE ONLY public.cart_items DROP CONSTRAINT cart_items_pkey;
       public                 raihan    false    226            �           2606    17992     order_details order_details_pkey 
   CONSTRAINT     ^   ALTER TABLE ONLY public.order_details
    ADD CONSTRAINT order_details_pkey PRIMARY KEY (id);
 J   ALTER TABLE ONLY public.order_details DROP CONSTRAINT order_details_pkey;
       public                 raihan    false    224            �           2606    17976    orders orders_pkey 
   CONSTRAINT     P   ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);
 <   ALTER TABLE ONLY public.orders DROP CONSTRAINT orders_pkey;
       public                 raihan    false    222            �           2606    17912    products products_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.products DROP CONSTRAINT products_pkey;
       public                 raihan    false    220            �           2606    17900    users users_email_key 
   CONSTRAINT     Q   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);
 ?   ALTER TABLE ONLY public.users DROP CONSTRAINT users_email_key;
       public                 raihan    false    218            �           2606    17896    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public                 raihan    false    218            �           2606    17898    users users_username_key 
   CONSTRAINT     W   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);
 B   ALTER TABLE ONLY public.users DROP CONSTRAINT users_username_key;
       public                 raihan    false    218            �           2606    18033    accounts accounts_user_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;
 H   ALTER TABLE ONLY public.accounts DROP CONSTRAINT accounts_user_id_fkey;
       public               raihan    false    3258    218    228            �           2606    18018 %   cart_items cart_items_product_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.cart_items
    ADD CONSTRAINT cart_items_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE CASCADE;
 O   ALTER TABLE ONLY public.cart_items DROP CONSTRAINT cart_items_product_id_fkey;
       public               raihan    false    226    3262    220            �           2606    18013 "   cart_items cart_items_user_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.cart_items
    ADD CONSTRAINT cart_items_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;
 L   ALTER TABLE ONLY public.cart_items DROP CONSTRAINT cart_items_user_id_fkey;
       public               raihan    false    226    3258    218            �           2606    17993 )   order_details order_details_order_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.order_details
    ADD CONSTRAINT order_details_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id) ON DELETE CASCADE;
 S   ALTER TABLE ONLY public.order_details DROP CONSTRAINT order_details_order_id_fkey;
       public               raihan    false    224    222    3264            �           2606    17998 +   order_details order_details_product_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.order_details
    ADD CONSTRAINT order_details_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id) ON DELETE CASCADE;
 U   ALTER TABLE ONLY public.order_details DROP CONSTRAINT order_details_product_id_fkey;
       public               raihan    false    220    3262    224            �           2606    17977    orders orders_user_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;
 D   ALTER TABLE ONLY public.orders DROP CONSTRAINT orders_user_id_fkey;
       public               raihan    false    218    3258    222            k   <   x���v
Q���W((M��L�KLN�/�+)Vs�	uV�0�Q "S=Mk... �[d      i   Y   x���v
Q���W((M��L�KN,*��,I�-Vs�	uV�0�Q0�Q0#u##S]#]##K+S#+cS=Sccu����\\\ ���      g   �   x�ŏA�@�����Y�.3�j�v�`AR�F�A(��%��[�x�����*�E�����wW}�v85����8fy�JX��X��ID�(غ]V畂���9bJ-��E'b�%���M��:��c��`����eԴfZq8M?A� K]�      e   �  x�řOo1�����vd;q�'=TBE���T*����O2�[mh,y����#�o�����9=;?y����}x�6?�\_]L�w��{���ۏ'��%��Ô�/.n�o��?֗/�?�!#8�y��+�z}t�,+��\.S)�6������ם;�C����(���,����(r�g�^y��W.�O��>�\�������8�z�!TݵzF�����oW��Rp��y���	�Z�i!�D��`��S]B�<åO�Z�!֯?��c,Zz����C?{��{�1��6a���3��V��7����Pa�A�E����0��~2Ï�c,^��/��g��Bi��+���҄�5Ƅu�[��v����jF�
!��1�1Q;�@3�s�%?su�4�Y��\����%,�:�[yy���$7�����;$���M0;i�J�d�͒u{ShF�,�c&�d��`���>����epHvղ�	�;�yHv޲=�>�8$�4��;�ؼ�}S�z��?̗o��N�ظ�=I<ӧ)Ŝ)<su��=S��Kf�g꒣�3u�I����,�L]vy�*;��3u�(�L]6�<S����K"��e��3u�Q����d��l���b�5�U��h�1�n��B7V��Y���.t�d�q6ݸX�} +1�h(����0z31��P#ۉa�Vb���l'����aC1Lh(�����zC1L�N��a��b����l%����a;1�h&���0{I��i�R��sqƃeM,beMn�ʚ�$��5�Y(kr�(NV$�ɚdDɚ\ɚd/��5�A"krY!k��(@�$'A|������O��s����W��e���0l��hvGu4	R�=R6�g�S���Y����/��Y`���2���e�G�S��T�Ӱ}*��m���$�c{Tm9P��L\�u�P]��ea��>����?a&L��Պ>:�D���      c   �  x����n�H��y��^2+`l���3�:ֆ2���vӁV��n�qN��z�$[��V�9�+P�_W�U?˻�zu˻��д%�Ect�rg�av��:��G8�5� �R��1�F���[�u��[�J�Sz���:'�e[�F�;
�8���0�1CV1�
��0�ñ�^@|��O���$�&���ϖ� RNVٶ.`�{�A���r��(VXB;|Լ�g���c}?I��JtQ��G���̅R�N3�-弑E�5�<iS��P�q��1��
�h��;�Q����,!��=j�w���a0#L�0~��3̤�$��]a=W��?��^��M��W����9'L'�p��:�.�z�7^���:�z����*��/��>�жŒKm�����LIG��=�@g�k���B�ׅt����!a#J~m
]�Jr4B�P�����������+a����SL�*��stn_V��'�`����j�gl;}<� ���yt�'5i�E�v�I������F�/+ eXM���p�Q�m�	`7��T���|9���E��7�[��VAaXSJnq�p�- P��;�͖49�c�)�\U�:��0m��d�_�b��m�2�[�C+]K��]��ȿy#���%,2��P������O}��	�EG��^CuFT�6J��V~oe+�d�`=_A.m��>�N#"��%�Zh] i*_p�_����_[l����ޠ�z���"	]*>]�C�T@.�b����.�G��-<D�|��S��B��L���m�A��a������2�Y2���Q�x�w �x��C�#�&��_9P�z�ɩ���p��:}dw���$5�������]�o�� V����8m{'%�h���{�Q0C��9��ٿ]-�W      a   �  x���]��0���W�.
nh���f)�,��.x�6��$��n�ߛ�[F0(���Rz^Nޜ�j���������V�A���6�]���~ �����xx��~3?��~�aR�z�n5�
5�?���7�!�&�c��YL�����W⬙���o����̬z�b����g��! ��X��))��E<����:I3���%q֍`FI���g���D ��Y��I�e�#J��ɜ5�y�{hR8��X�%��聡�,�c/:�w��J�5+P�#��%\�Zj}TK����"��\Z��~�6�M?�Z�K����&R�;D<6%ך�և5G��hd���ƺ�bXw��G5��#-��&�X�_sWQ�����t�����A��)xg���U�EG�b�
k�PiY͕b��OM���A�K�����Ț�x+5��:��^���:�V? �Ī�     