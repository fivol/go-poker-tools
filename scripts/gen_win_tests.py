
source = """
1.2s3s4s5s6s Ac5c AsQs Ad2h, Ac5c AsQs Ad2h

2.2s2d2h2c3c Ac9d QsQd Ad3d, Ac9d Ad3d

3.5d5h5c3c3h 2h4d AsKh 2c2d, 2h4d AsKh 2c2d

4.5d5h5c3c3h 2h4d AsKh 4c4d, 4c4d

5.6s6d6c2c3h AsKd AcQh AhJh, AsKd

6.6s6d6cAhKc QhJc QcTc 2h3h, QhJc QcTc 2h3h

7.2c2h3s3dAc KhQh Qc5c 4h6h, KhQh Qc5c 4h6h

8.2c2h3s3dAc KhQh Qc5c 4h4c, 4h4c

9.KcKhQcQd2s 3s3d 5s5d 7s6d, 7s6d

10.KcKhQcQd8s 3s3d 5s5d 7s6d, 3s3d 5s5d 7s6d

11.AhKhQhJh8h 7h6d 5s5d KcQc, 7h6d 5s5d KcQc

12.AhKhQhJh6h 7h6d 5s5d KcQc, 7h6d 

13.8s7c6d5h4c AcAd KcKd QcQd, AcAd KcKd QcQd

14.8s7c6d5h4c AcAd KcKd 9c9d, 9c9d

15.8c8d7d6d2c 8s8h Td9d 7c7s, Td9d

16.AcJsTd9hKc Ah7c As2c Ks8d, Ah7c As2c

17.Ac9sTd8hKc AhQc As2c Ks2d, AhQc

18.KcKdKh4c3c 4h4d 3h3d 5h5d, 5h5d

19.Ac2d3h4c5d AhAd 6s9d 6c7d, 6c7d

20.AcAdAhAsKs KcKd 2s3d 8s8d, KcKd 2s3d 8s8d

"""

if __name__ == '__main__':
    lines = source.split('\n')
    lines = list(filter(lambda x: x.strip(), lines))
    for line in lines:
        line = line.split('.')[1]
        # print(line)
        board = line.split(' ')[0]
        hands = line.split(',')[0].split(' ')[1:]
        winner_hands = line.split(',')[1].strip().split(' ')
        # print(board, hands, winner_hands)
        winner_hands = [hands.index(hand) for hand in winner_hands]

        template = f"""{{
    "{board}",
    []string{{{', '.join(map(lambda x: f'"{x}"', hands))}}},
    []int{{{', '.join(map(str, winner_hands))}}},
}},"""
        print(template)
