BEGIN;

INSERT INTO role (title) VALUES 
('admin'),
('user'),
('moderator'),
('writter');

INSERT INTO users (username, email, password, image, role_id) VALUES 
('gizmo', 'gizmo@mail.com', '$2a$10$KNmL3aNFAgmlVTH9.937HehS2lfUVgRp1lC4aai0bC2eOAAYB8TDO', 'https://randomuser.me/api/portraits/men/9.jpg', '1'),
('Alexandrea', 'Alexandrea@mail.com', '$2a$10$KNmL3aNFAgmlVTH9.937HehS2lfUVgRp1lC4aai0bC2eOAAYB8TDO', 'https://randomuser.me/api/portraits/men/8.jpg', '1'),
('Alyce', 'Alyce@mail.com', '$2a$10$KNmL3aNFAgmlVTH9.937HehS2lfUVgRp1lC4aai0bC2eOAAYB8TDO', 'https://randomuser.me/api/portraits/men/8.jpg', '2'),
('Alden', 'Alden@mail.com', '$2a$10$KNmL3aNFAgmlVTH9.937HehS2lfUVgRp1lC4aai0bC2eOAAYB8TDO', 'https://randomuser.me/api/portraits/men/6.jpg', '1'),
('Kaci', 'Kaci@mail.com', '$2a$10$KNmL3aNFAgmlVTH9.937HehS2lfUVgRp1lC4aai0bC2eOAAYB8TDO', 'https://randomuser.me/api/portraits/men/10.jpg', '2'),
('Johan', 'Johan@mail.com', '$2a$10$KNmL3aNFAgmlVTH9.937HehS2lfUVgRp1lC4aai0bC2eOAAYB8TDO', 'https://randomuser.me/api/portraits/men/1.jpg', '1');

INSERT INTO category (title) VALUES 
('salade'),
('fast food'),
('vegan'),
('viande'),
('desert'),
('entree'),
('plat'),
('gouter');

INSERT INTO recipe (title, describ, time_prep, image, user_id, category_id) VALUES
('pariaturut', 'Officiis illo maxime sit ut est. Perspiciatis voluptatem earum id quas culpa. Itaque incidunt qui et quos tenetur. Pariatur error non aspernatur iusto in. Sunt qui quae rerum qui porro. Explicabo tenetur esse et ea voluptatum. Quidem sed voluptatem voluptas ratione doloremque. Minima nostrum molestias iusto omnis molestiae.', '128', 'https://picsum.photos/200/300?random=248', '1', '3'),
('laboreharum', 'Dolorum ut enim temporibus molestiae perspiciatis.', '47', 'https://picsum.photos/200/300?random=839', '1', '7'),
('excepturicupiditate', 'Accusantium harum explicabo eveniet nam tenetur. Eos voluptatem quaerat unde vitae quibusdam. Voluptatibus assumenda eos illum ut sed. Consequatur autem rem debitis officia autem. Nihil at nostrum beatae numquam quia. Sapiente inventore voluptatem qui quis consectetur.', '76', 'https://picsum.photos/200/300?random=834', '1', '3'),
('utquasi', 'Eos eveniet eum voluptas architecto eligendi. Qui nostrum unde officiis nihil rerum. Non vel qui sed doloremque velit. Inventore perferendis aut saepe eius laborum. Consequatur est eligendi laudantium nemo aliquid. Neque perspiciatis quae velit beatae facere.', '111', 'https://picsum.photos/200/300?random=973', '1', '1'),
('eaet', 'Nemo minus dolores ut dolorem non. Dolorem vel exercitationem consequatur laborum at. Sed et quae ad in ratione.', '190', 'https://picsum.photos/200/300?random=831', '2', '6'),
('perferendisvel', 'Id nostrum et voluptatem facilis assumenda. Aliquid quas porro voluptatum est sed. Id quo minima ad totam porro. Id enim architecto earum repellendus assumenda. Et et enim dolorem cum quam. Et aliquam rerum est excepturi aut. Laudantium et nemo doloribus et labore.', '54', 'https://picsum.photos/200/300?random=583', '3', '3'),
('quibusdamaliquam', 'Exercitationem vitae atque possimus nemo distinctio. Quia inventore reiciendis est est molestiae. Voluptatem nisi blanditiis et debitis voluptatibus. Rerum vero quae iusto voluptatem consequuntur. Enim perspiciatis ipsum sunt in neque. Eum qui aliquid non qui qui. Libero voluptatum voluptas et et sit. Magnam nesciunt nobis deleniti ullam repellat. Accusantium nostrum repudiandae consequatur asperiores qui.', '94', 'https://picsum.photos/200/300?random=248', '2', '6'),
('maioresdistinctio', 'Aut esse inventore qui quod impedit. Eligendi dolores in est inventore sit. Adipisci distinctio fuga et fugit aut. Voluptate dolorem dolorem vitae est iste.', '43', 'https://picsum.photos/200/300?random=345', '1', '7'),
('etea', 'Enim odit mollitia omnis possimus commodi. Expedita ut adipisci aut esse sunt.', '38', 'https://picsum.photos/200/300?random=426', '1', '3'),
('voluptatenumquam', 'Ratione dolor quia porro et tenetur. Omnis blanditiis quo aliquid quae temporibus. A est sit dolorum possimus quos. Quasi totam omnis veniam qui voluptate. Molestias explicabo eos consectetur cum sequi. Nesciunt repellendus voluptatem recusandae et sed. Iure ea et rerum est beatae.', '160', 'https://picsum.photos/200/300?random=382', '1', '3'),
('quiavelit', 'Labore natus quod aspernatur exercitationem quia. Aliquid corrupti dignissimos natus ipsam odit. Nemo consectetur voluptas et sed voluptatem. Ipsa qui molestiae quae dolor similique. Nam occaecati vel autem modi saepe. Unde et harum architecto maiores cupiditate. Tempora veniam fugiat voluptatem tenetur qui. In ut eius eveniet ut excepturi. Optio sit et quos quibusdam ab. Consequatur libero quasi aut et nostrum.', '72', 'https://picsum.photos/200/300?random=978', '2', '4'),
('omnisad', 'Exercitationem perferendis natus velit dolores voluptas. Commodi sunt culpa beatae consequatur veritatis. Maxime non nihil ut est ut. Quae animi soluta sed fugit provident. Reprehenderit cum iste sint vero mollitia.', '137', 'https://picsum.photos/200/300?random=566', '1', '3'),
('etdebitis', 'Ea tenetur dolore voluptatem nobis aut. Nostrum consequuntur vero rerum atque dolor. Voluptas officiis qui hic laborum et. Et cum totam et amet vel. Repellat et maxime molestias voluptate ut. Harum sint nulla repudiandae qui rerum. Earum omnis ducimus ut cumque laudantium.', '63', 'https://picsum.photos/200/300?random=86', '1', '1'),
('cumquae', 'Tempora qui incidunt temporibus qui neque. Officiis iste tempora et incidunt asperiores. Ipsum ut omnis nobis rerum nisi. Quod nemo dolorem necessitatibus accusantium qui. Dolor dolor tempora aliquid voluptatem quia. Dolorem debitis quae fuga officiis ut. Deserunt aut quis id rem numquam. Assumenda explicabo nihil harum labore officia. Voluptatum fuga laborum nobis atque sed.', '100', 'https://picsum.photos/200/300?random=693', '1', '4'),
('etesse', 'Aut quia et saepe quos atque. Exercitationem praesentium voluptatibus dicta adipisci incidunt. Qui suscipit consequuntur animi at dolor. Necessitatibus distinctio omnis aut rerum quis. Accusamus molestiae autem quis perspiciatis iste. Vero doloribus tenetur possimus dolor est. Laboriosam quia sequi repellat necessitatibus harum.', '119', 'https://picsum.photos/200/300?random=168', '2', '2'),
('quoveritatis', 'Repudiandae commodi labore quibusdam dolorum quas. Harum quibusdam aut doloremque recusandae at. Ut eos voluptatem dicta aut doloribus. Pariatur rem quis libero enim ut. Maxime aut enim omnis quisquam tempore. Omnis magni quod quasi consequatur quia. Totam saepe facere molestias perferendis vel.', '41', 'https://picsum.photos/200/300?random=58', '2', '2'),
('debitisiusto', 'Aliquid voluptatum qui quis architecto ut. Placeat facere sit et saepe vel. Adipisci est facere consequatur sapiente nam. Qui placeat debitis dicta blanditiis est. Minus sint incidunt quae est consequatur. Ratione enim autem cupiditate nobis debitis.', '19', 'https://picsum.photos/200/300?random=698', '1', '4'),
('iureut', 'Neque quidem cum fugiat et quam. Blanditiis eum temporibus vel qui et. Delectus ut ipsa quas modi facilis. Magni temporibus odit quia est illo. Eius sint aut est et consequatur. Quam qui vel enim dolores sint.', '176', 'https://picsum.photos/200/300?random=678', '1', '1'),
('omnisnisi', 'Delectus corporis sit et impedit numquam. Est vel soluta tempora ut aut.', '111', 'https://picsum.photos/200/300?random=854', '2', '2'),
('doloremquelaudantium', 'Est nihil ducimus vel quidem repellendus. Qui velit ab est et accusamus. Expedita tempora enim natus saepe ipsa. Unde sit facere consequatur vel est. Qui ducimus architecto eaque quae pariatur. Odit est praesentium dolores sunt repudiandae. Totam dolor a corrupti inventore dicta. Autem debitis incidunt nam sit distinctio. Odit sequi provident qui itaque eum. Temporibus hic facere sint qui tempore.', '157', 'https://picsum.photos/200/300?random=654', '1', '4'),
('inaliquam', 'Reiciendis qui quidem velit consequatur eligendi. Qui quibusdam expedita provident et nobis. Temporibus corporis est dolorum dolorem occaecati. Id ab quae autem repudiandae ea. Iste et animi mollitia quia nemo. Cum laboriosam eligendi dicta quia consectetur. Voluptas quas natus ipsa cum omnis.', '195', 'https://picsum.photos/200/300?random=512', '1', '2'),
('eaqui', 'Accusamus tenetur autem est sed aut. Qui delectus autem ducimus sed commodi. Est minima qui rem blanditiis impedit. Consectetur sed perspiciatis in earum odio. Qui cupiditate inventore dolorem aperiam quo. Ratione eum officiis velit quas odio. Perferendis eos molestias vero eveniet omnis. Autem qui recusandae quos commodi dolores. Minima quia est dolorem nisi voluptate. Iste et eum et quidem nisi.', '186', 'https://picsum.photos/200/300?random=962', '3', '4'),
('doloremfugiat', 'Similique consequuntur aut ratione magnam aliquid. Voluptatum hic est et odio quo. Mollitia iure alias inventore ratione dolores. Recusandae et a tempora nihil omnis. Vel quia et et eum voluptatum.', '101', 'https://picsum.photos/200/300?random=145', '2', '6'),
('suscipitqui', 'Accusamus illum quos minus assumenda voluptatem. Voluptate repellendus labore vel praesentium sint. Harum sapiente excepturi ratione atque et. Illo quas dolor saepe sed cumque. Quo dolores nobis et magnam fuga. Illum necessitatibus magnam qui aspernatur similique. Et magnam aut vel blanditiis impedit. Aut atque animi numquam id distinctio. Nam quis nulla modi esse quia. Laudantium sequi et nulla dolorem omnis.', '84', 'https://picsum.photos/200/300?random=281', '1', '7'),
('automnis', 'Dolore sunt tenetur amet illo fuga. Qui temporibus dolores cum optio excepturi. Cupiditate dolorem aut modi fugiat voluptate. Laboriosam quia molestiae enim incidunt quidem. Voluptates earum impedit error qui labore. Tempore id ut animi dolor dolore. Voluptatum aliquid commodi officia quam est.', '2', 'https://picsum.photos/200/300?random=816', '1', '3'),
('corporisvoluptatibus', 'Nemo nostrum nulla explicabo ea officiis. Cupiditate deserunt at veniam iure provident. Velit et quod accusantium illum est. Veritatis consequatur sit necessitatibus voluptates nihil. Assumenda quaerat reprehenderit sapiente nemo nihil. Corporis qui nemo rerum voluptate quo.', '124', 'https://picsum.photos/200/300?random=419', '2', '2'),
('mollitiadolorum', 'Qui reiciendis nemo consequatur dolores vel. Non enim est earum dolor nesciunt. Et est beatae aut accusantium magni. Accusamus maiores laborum rem eos quo. Autem necessitatibus non exercitationem et quo. Alias harum sint nemo ut amet. Fuga illo architecto autem aliquid quo.', '40', 'https://picsum.photos/200/300?random=277', '3', '4'),
('quivoluptas', 'Sunt eum expedita suscipit nihil maxime. Odio et et iure delectus adipisci. Nam minus laborum et quibusdam vero. Hic qui et ea aut maxime. Deleniti iure ut voluptatem quas magni. Natus sint pariatur tenetur accusantium quo. Nisi recusandae debitis est expedita numquam.', '104', 'https://picsum.photos/200/300?random=589', '1', '6'),
('quaeratqui', 'Consequuntur consequatur qui excepturi facere repudiandae.', '109', 'https://picsum.photos/200/300?random=176', '1', '5'),
('ateius', 'Quam necessitatibus error eos iure soluta. Repudiandae animi amet sed est doloribus. Corrupti voluptates omnis consectetur ducimus vitae. Pariatur occaecati aspernatur omnis sit dolorem. Recusandae rem aut earum rerum reiciendis. Est sit odio qui recusandae velit. Cumque veniam quod qui consequatur placeat.', '176', 'https://picsum.photos/200/300?random=542', '1', '4');

COMMIT;